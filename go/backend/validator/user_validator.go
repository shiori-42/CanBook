package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/shiori-42/textbook_change_app/go/backend/model"
)

func UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.Length(1, 255).Error("email must be between 1 and 255 characters"),
			is.Email.Error("invalid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.Length(6, 255).Error("password must be between 6 and 255 characters"),
		),
	)
}
