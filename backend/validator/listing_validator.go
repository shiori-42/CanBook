package validator

import (
	"github.com/shiori-42/textbook_change_app/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IListingValidator interface {
	ListingValidate(listing model.Listing) error
}

type listingValidator struct{}

func NewListingValidator() IListingValidator {
	return &listingValidator{}
}

func (lv *listingValidator) ListingValidate(listing model.Listing) error {
	return validation.ValidateStruct(&listing,
		validation.Field(
			&listing.BookTitle, 
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 10).Error("limited max 10 characters"),
		),
	)
}