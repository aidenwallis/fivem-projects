package main

import (
	"context"
	"flag"
	"fmt"
	golog "log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/backend"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/config"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/expiry"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/migrations"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/privateapi"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/publicapi"
	"github.com/aidenwallis/go-utils/utils"
	"github.com/pkg/errors"
	"github.com/uptrace/bun/migrate"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var (
	configFile = flag.String("config", "./config.json", "The location of the config file.")
)

func main() {
	// flag.Usage = help
	flag.Parse()
	cmd := flag.Arg(0)

	if configFile == nil || *configFile == "" {
		golog.Fatalln("Cannot find valid config file location, please specify with --config=./config.json")
	}

	cfg, err := config.NewAppConfig(*configFile)
	if err != nil {
		golog.Fatalln(err.Error())
	}

	log, err := config.NewLogger(cfg.Environment)
	if err != nil {
		golog.Fatalln("Failed to init logger: " + err.Error())
	}

	dbImpl, err := db.NewDB(cfg.Database, log)
	if err != nil {
		log.Fatal("Failed to connect to database", zap.Error(err))
	}

	switch cmd {
	case "start", "run":
		start(cfg, dbImpl, log)

	case "make-migration":
		makeMigration(dbImpl)

	case "rollback":
		rollback(dbImpl, log)

	default:
		help()
	}
}

func help() {
	golog.Println("aiden_auth <command>")
	golog.Println("  run            Run the service")
	golog.Println("  make-migration Create a new migration")
	golog.Println("  rollback       Rollback previous migration")
}

func makeMigration(dbImpl db.DB) {
	name := flag.Arg(1)
	if name == "" {
		golog.Println("usage: aiden_auth make-migration <name_of_migration>")
		return
	}
	migrator := migrate.NewMigrator(dbImpl.Bun(), migrations.Migrations)
	_, _ = migrator.CreateGoMigration(context.Background(), name)
	golog.Println("Done.")
}

func rollback(dbImpl db.DB, log config.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	group, err := migrate.NewMigrator(dbImpl.Bun(), migrations.Migrations).Rollback(ctx)
	if err != nil {
		log.Fatal("Failed to rollback", zap.Error(err))
		return
	}

	log.Info(utils.Ternary(group.IsZero(), "Nothing to rollback...", fmt.Sprintf("Rolled back %s", group)))
}

func start(cfg *config.AppConfig, dbImpl db.DB, log config.Logger) {
	if err := doMigrate(dbImpl, log); err != nil {
		log.Fatal("Failed to assert database version", zap.Error(err))
	}

	backendImpl := backend.NewBackend(dbImpl, log, cfg.Sessions)

	cancelPublic, err := serverFactory("public", cfg.Servers.Public, log, publicapi.NewServer(backendImpl, log))
	if err != nil {
		log.Fatal("Failed to start public server", zap.Error(err))
	}

	cancelPrivate, err := serverFactory("private", cfg.Servers.Private, log, privateapi.NewServer(backendImpl, log))
	if err != nil {
		log.Fatal("Failed to start private server", zap.Error(err))
	}

	// Start expiry worker to evict old tokens from db at a fixed frequency
	go expiry.NewWorker(dbImpl, log).Start()

	log.Info("Service successfully started. Press CTRL+C to exit.")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	log.Info("Shutting down...")

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return errors.Wrap(cancelPublic(ctx), "shutting down public api")
	})

	g.Go(func() error {
		return errors.Wrap(cancelPrivate(ctx), "shutting down private api")
	})

	if err := g.Wait(); err != nil {
		log.Error("Failed to shut down all systems.", zap.Error(err))
	}

	log.Info("Goodbye!", zap.Int64("shutdown-ms", time.Since(start).Milliseconds()))
}

func doMigrate(dbImpl db.DB, log config.Logger) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	log = log.With(zap.String("task", "db-migrations"))

	migrator := migrate.NewMigrator(dbImpl.Bun(), migrations.Migrations)

	log.Info("Migrating database...")

	if err := migrator.Init(ctx); err != nil {
		log.Warn("Skipping init step...", zap.Error(err))
	}

	group, err := migrator.Migrate(ctx)
	if err != nil {
		return err
	}

	log.Info(utils.Ternary(group.IsZero(), "No new migrations, skipping...", fmt.Sprintf("Migrated to %s!", group)))

	return nil
}

func serverFactory(name string, cfg *config.ServerConfig, log config.Logger, handler http.Handler) (func(context.Context) error, error) {
	srv := &http.Server{
		Handler: handler,
	}

	log = log.With(zap.String("server-name", name))

	ln, err := net.Listen(cfg.Transport, cfg.Addr)
	if err != nil {
		return nil, errors.Wrapf(err, "listening for %s", name)
	}

	log.Info("Server started.", zap.String("transport", cfg.Transport), zap.String("addr", cfg.Addr))

	ch := make(chan struct{}, 1)

	go func(srv *http.Server) {
		err := srv.Serve(ln)
		select {
		case <-ch:
			// deliberate exit, ignore
			return
		default:
			log.Error("failed to start server", zap.Error(err))
		}
	}(srv)

	return func(ctx context.Context) error {
		close(ch)
		return srv.Shutdown(ctx)
	}, nil
}
