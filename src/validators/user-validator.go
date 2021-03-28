package validators

import (
	"fmt"

	"github.com/challenge-bw-go/src/models"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func UserCreateValidator(user models.User) bool {
	validate = validator.New()
	err := validate.Struct(user)

	if len(user.Username) < 5 || len(user.Username) > 30 {
		return false
	}

	if len(user.Name) < 3 || len(user.Name) > 30 {
		return false
	}

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return false
		}

		for _, err := range err.(validator.ValidationErrors) {

			if err != nil {
				fmt.Println(err)             // can differ when a custom TagNameFunc is registered or
				fmt.Println(err.Namespace()) // can differ when a custom TagNameFunc is registered or
				fmt.Println(err.Field())     // by passing alt name to ReportError like below
				fmt.Println(err.Param())
				fmt.Println()
				return false
			}

		}
	}

	return true
}
