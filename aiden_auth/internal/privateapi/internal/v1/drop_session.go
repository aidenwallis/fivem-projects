package v1

import (
	"net/http"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/schema"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/utils"
	"github.com/aidenwallis/go-write/write"
	"go.uber.org/zap"
)

func (v *Version) DropSession(w http.ResponseWriter, req *http.Request) {
	body, schemaErr := utils.ParseBody[schema.DropSessionInput](req)
	if schemaErr != nil {
		_ = write.BadRequest(w).JSON(schemaErr)
		return
	}

	if err := v.b.DropSession(req.Context(), body.Identifiers); err != nil {
		utils.WithRequestLogger(v.log, req).Error("failed to drop sessions", zap.Error(err))
		_ = write.InternalServerError(w).JSON(schema.UnknownError)
		return
	}

	_ = write.OK(w).Empty()
}
