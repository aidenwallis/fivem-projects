package expiry

import (
	"context"
	"time"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/db"
	"go.uber.org/zap"
)

// expireFrequency is how often we should poll the db and evict tokens
const expireFrequency = time.Minute * 5

type Worker struct {
	ctx       context.Context
	cancelCtx context.CancelFunc
	db        db.DB
	log       *zap.Logger
}

// NewWorker creates a new instance of worker
func NewWorker(dbImpl db.DB, log *zap.Logger) *Worker {
	ctx, cancel := context.WithCancel(context.Background())
	return &Worker{
		ctx:       ctx,
		cancelCtx: cancel,
		db:        dbImpl,
		log:       log,
	}
}

// Start opens the goroutine that will wake up and evict at a fixed frequency
func (w *Worker) Start() {
	for {
		if err := w.tick(); err != nil {
			w.log.Error("failed to expire sessions", zap.Error(err))
		}

		select {
		case <-w.ctx.Done():
			// loop ended, exit goroutine
			return

		case <-time.After(expireFrequency):
			// due another eviction, unblock goroutine
			continue
		}
	}
}

func (w *Worker) tick() error {
	ctx, cancel := context.WithTimeout(w.ctx, time.Minute)
	defer cancel()

	start := time.Now().UTC()
	w.log.Info("expiring sessions...")

	count, err := w.db.ExpireSessions(ctx)
	if err != nil {
		return err
	}

	w.log.Info(
		"finished expiring sessions",
		zap.Int("deleted-count", count),
		zap.Int64("duration-ms", time.Since(start).Milliseconds()),
	)

	return nil
}
