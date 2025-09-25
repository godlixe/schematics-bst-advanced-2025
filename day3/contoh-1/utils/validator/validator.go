package validatorx

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ValidationErrorItem struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

func ParseValidatorErrors(err error) ([]ValidationErrorItem, error) {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return nil, err
	}

	out := make([]ValidationErrorItem, 0, len(ve))
	for _, fe := range ve {
		out = append(out, ValidationErrorItem{
			Field:   fe.Field(),
			Tag:     fe.Tag(),
			Message: fe.Error(),
		})
	}
	return out, nil
}
