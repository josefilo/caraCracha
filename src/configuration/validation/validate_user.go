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

var translator ut.Translator

func init() {
	if validador, ok := binding.Validator.Engine().(*validator.Validate); ok {
		englishLanguage := en.New()
		universalTranslator := ut.New(englishLanguage, englishLanguage)
		translator, _ = universalTranslator.GetTranslator("en")
		enTranslation.RegisterDefaultTranslations(validador, translator)
	}
}

func ValidationUserError(validationErr error) *api_errors.ApiError {
	var unMarshalTypeError *json.UnmarshalTypeError
	var ValidationErrors validator.ValidationErrors

	if errors.As(validationErr, &unMarshalTypeError) {
		return api_errors.NewBadRequestError("Invalid field type")
	} else if errors.As(validationErr, &ValidationErrors) {
		errorsCauses := []api_errors.Cause{}
		for _, errors := range validationErr.(validator.ValidationErrors) {
			cause := api_errors.Cause{
				Message: errors.Translate(translator),
				Field:   errors.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return api_errors.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return api_errors.NewBadRequestError("Error trying to convert fields")
	}
}
