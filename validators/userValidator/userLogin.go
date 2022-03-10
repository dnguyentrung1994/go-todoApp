package userValidators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateUserLogin(data interface{}) []string {
	validate = validator.New()

	if loginAccount, ok := data.(LoginForm); ok {
		err := validate.Struct(loginAccount)
		var output []string

		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				output = append(output, "InvalidValidator!")
			}
			for _, err := range err.(validator.ValidationErrors) {
				output = append(output, fmt.Sprintf("NameSpace %s: ErrorTag %s", err.Namespace(), err.Tag()))
			}
		}

		return output
	}
	return []string{"VALIDATOR_MISMATCH"}
}
