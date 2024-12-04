package utils

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/MatthewAraujo/min-ecommerce/types"
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return errors.New("request body is empty")
	}
	return json.NewDecoder(r.Body).Decode(payload)

}

func ToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func TranslateValidationErrors(errs validator.ValidationErrors) []types.ValidationErrorResponse {
	var errors []types.ValidationErrorResponse
	for _, err := range errs {
		message := fmt.Sprintf("O campo '%s' falhou na validação: %s", err.Field(), err.Tag())
		if err.Tag() == "required" {
			message = fmt.Sprintf("O campo '%s' é obrigatório.", err.Field())
		} else if err.Tag() == "oneof" {
			message = fmt.Sprintf("O campo '%s' deve ser um dos seguintes valores: %s.", err.Field(), err.Param())
		}
		errors = append(errors, types.ValidationErrorResponse{
			Field:      err.Field(),
			Validation: err.Tag(),
			Value:      err.Param(),
			Message:    message,
		})
	}
	return errors
}
