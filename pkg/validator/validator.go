package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Validate(s interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(s)

	if err == nil {
		return nil
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errorMessages := make([]string, 0)
		for _, fieldErr := range validationErrors {
			msg := fmt.Sprintf("error in field %s:%s",
				fieldErr.Field(),
				fieldErr.Tag())
			errorMessages = append(errorMessages, msg)
		}
		return fmt.Errorf(strings.Join(errorMessages, ";"))
	}
	return fmt.Errorf("unknow validator error")
}
