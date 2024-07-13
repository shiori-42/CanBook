package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/shiori-42/textbook_change_app/go/backend/model"
)

func ItemValidate(item model.Item) error {
	return validation.ValidateStruct(&item,
		validation.Field(
			&item.TextName,
			validation.Required.Error("Textname is required"),
			validation.Length(1, 255).Error("Textname must be between 1 and 255 characters"),
		),
	)
}
