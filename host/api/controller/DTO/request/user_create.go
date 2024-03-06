package request

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type UserCreate struct {
	Name     string `validate:"required,min=4,max=30"`
	Password string `validate:"required,customPass,min=8,max=30"`
	Email    string `validate:"required,email"`
}

func customPass(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	//(UPPER)(number)(special c)
	regex := `[A-Za-z].*\d.*[@$!%*?&].*`

	match, err := regexp.MatchString(regex, password)
	if err != nil {
		return false
	}
	return match
}
