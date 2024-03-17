package validator

import (
	"github.com/shiori-42/textbook_change_app/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IListingValidator interface {
	ListingBalidate(listing *model.Listing) error
}

type listingValidator struct {}

func NewListingValidator() IListingValidator {
	return &listingValidator{}
}

func (lv *listingValidator) ListingBalidate(listing *model.Listing) error {
	return validation.ValidationStruct(&listing,
		validation.Field(
			&listing.Title, 
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 10).Error("limited max 10 characters")
		),
	)
}