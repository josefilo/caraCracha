package validation

import (
	"encoding/json"
	"errors"

	"github.com/Josefreitas788/gestao_de_financas/src/configuration/api_errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
)

var (
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)
		transl, _ = un.GetTranslator("en")
		enTranslation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidationUserError(validationErr error) *api_errors.ApiError {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationErr validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return api_errors.NewBadRequestError("Invalid field type")
	} else if errors.As(validationErr, &jsonValidationErr) {
		errorsCauses := []api_errors.Cause{}
		for _, errors := range validationErr.(validator.ValidationErrors) {
			cause := api_errors.Cause{
				Message: errors.Translate(transl),
				Field:   errors.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return api_errors.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return api_errors.NewBadRequestError("Error trying to convert fields")
	}
}
