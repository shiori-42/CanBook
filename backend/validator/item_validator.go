package validator

import (
	"github.com/shiori-42/textbook_change_app/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ItemValidate(item model.Item) error {
	return validation.ValidateStruct(&item,
		validation.Field(
			&item.Name,
			validation.Required.Error("name is required"),
			validation.Length(1, 255).Error("name must be between 1 and 255 characters"),
		),
		validation.Field(
			&item.CategoryID,
			validation.Required.Error("category_id is required"),
		),
	)
}