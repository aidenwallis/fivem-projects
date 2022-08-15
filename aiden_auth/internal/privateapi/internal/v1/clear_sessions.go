package v1

import (
	"net/http"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/schema"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/utils"
	"github.com/aidenwallis/go-write/write"
	"go.uber.org/zap"
)

// ClearSessions will revoke all sessions in the DB.
func (v *Version) ClearSessions(w http.ResponseWriter, req *http.Request) {
	if err := v.b.ClearSessions(req.Context()); err != nil {
		utils.WithRequestLogger(v.log, req).Error("failed to clear sessions", zap.Error(err))
		_ = write.InternalServerError(w).JSON(schema.UnknownError)
		return
	}

	_ = write.OK(w).Empty()
}
