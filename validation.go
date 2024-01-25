package rest

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	enl := en.New()
	unt := ut.New(enl, enl)
	transl, _ = unt.GetTranslator("en")

	en_translation.RegisterDefaultTranslations(Validate, transl)
}

// Validate one struct error.
// Returns rest.Err with the cause of the error occurrence.
func ValidateStructErr(err error) *Err {
	var json_err *json.UnmarshalTypeError
	var validation_err validator.ValidationErrors

	if errors.As(err, &json_err) {
		return NewBadRequestErr("invalid field type", nil)
	} else if errors.As(err, &validation_err) {
		causes := []Cause{}

		for _, e := range err.(validator.ValidationErrors) {
			cause := Cause{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			causes = append(causes, cause)
		}

		return NewBadRequestErr("some field are invalid", causes)
	} else {
		return NewBadRequestErr("error trying to convert fields", nil)
	}
}

// Validate one variable error.
// Returns rest.Err with the cause of the error occurrence.
func ValidateVarErr(name string, err error) *Err {
	var validation_err validator.ValidationErrors

	if errors.As(err, &validation_err) {
		causes := []Cause{}

		for _, e := range err.(validator.ValidationErrors) {
			cause := Cause{
				Message: fmt.Sprintf("%s%s", name, e.Translate(transl)),
				Field:   name,
			}

			causes = append(causes, cause)
		}

		return NewBadRequestErr("some field are invalid", causes)
	} else {
		return NewBadRequestErr("error trying to convert fields", nil)
	}
}
