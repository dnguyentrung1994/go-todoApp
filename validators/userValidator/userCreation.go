package userValidators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateUserCreation(data interface{}) []string {
	validate = validator.New()

	if user, ok := data.(CreateUser); ok {
		err := validate.Struct(user)
		var output []string

		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				output = append(output, "InvalidValidator!")
				return output
			}

			for _, err := range err.(validator.ValidationErrors) {
				output = append(output, fmt.Sprintf("NameSpace %s: ErrorTag %s", err.Namespace(), err.Tag()))
			}
			return output
		}
		return output
	}
	return []string{"VALIDATOR_MISMATCH"}
}
