package v1

import (
	"net/http"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/schema"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/utils"
	"github.com/aidenwallis/go-write/write"
	"go.uber.org/zap"
)

func (v *Version) CreateSession(w http.ResponseWriter, req *http.Request) {
	body, schemaError := utils.ParseBody[schema.SessionInput](req)
	if schemaError != nil {
		_ = write.BadRequest(w).JSON(schemaError)
		return
	}

	sess, token, err := v.b.CreateSession(req.Context(), body.Identifiers, body.Metadata)
	if err != nil {
		utils.WithRequestLogger(v.log, req).Error("failed to create session", zap.Error(err))
		_ = write.InternalServerError(w).JSON(schema.UnknownError)
		return
	}

	_ = write.OK(w).JSON(schema.NewCreateSessionResponse(sess, token))
}
