package userValidators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateUserCreation(user *CreateUser) []string {
	validate = validator.New()

	err := validate.Struct(user)
	var output []string

	switch user.Role {
	case "GUEST":
	case "TEAM_MEMBER":
	case "TEAM_LEADER":
	case "MODERATOR":
	case "ADMIN":
		break
	default:
		output = append(output, "InvalidRole")
	}
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
