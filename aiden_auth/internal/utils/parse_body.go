package utils

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/schema"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/schema/codes"
	"github.com/aidenwallis/go-utils/utils"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
}

// ValidationError is how we format validation errors to end clients
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ParseBody[T any](req *http.Request) (*T, *schema.Error) {
	var body T
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		return nil, schema.InvalidBodyError
	}

	err := validate.Struct(&body)
	if err == nil {
		return &body, nil
	}

	if e, ok := err.(validator.ValidationErrors); ok {
		return nil, schema.NewError(
			codes.ValidationError,
			"Failed to validate body",
			utils.SliceMap(e, func(v validator.FieldError) *ValidationError {
				return &ValidationError{
					Field:   v.Field(),
					Message: v.Error(),
				}
			}),
		)
	}

	return nil, schema.NewError(codes.ValidationError, "Failed to validate body")
}
