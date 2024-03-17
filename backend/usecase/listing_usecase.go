package usecase

import (
	
	"github.com/shiori-42/textbook_change_app/model"
	"github.com/shiori-42/textbook_change_app/repository"
)

type IListingUsecase interface {
	GetAllMyListings(userId uint) ([]model.ListingResponse,error)
	GetMyListingById(userId uint, listingId uint) (model.ListingResponse,error)
	CreateListing(listing model.Listing) (model.ListingResponse,error)
	UpdateListing(listing model.Listing, userId uint, listingId uint)(model.ListingResponse,error)
	DeleteListing(userId uint, listingId uint) error
}

type listingUsecase struct {
	lr repository.IListingRepository
	tv validator.IListingValidator
}

func NewListingUsecase(lr repository.IListingRepository,lv validator.IListingValidator) IListingUsecase {
	return &listingUsecase{lr,lv}
}

func (lu *listingUsecase) GetAllMyListings(userId uint) ([]model.ListingResponse,error) {
	listings:=[]model.Listing{}
	if err := lu.lr.GetAllMyListings(&listings, userId); err != nil {
		return nil,err
	}
	resListings:=[]model.ListingResponse{}
	for _,v:=range listings{
		l:=model.ListingResponse{
			ID: v.ID,
			BookTitle: v.BookTitle,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resListings=append(resListings,l)
	}
	return resListings,nil
}

func (lu *listingUsecase) GetMyListingById(userId uint, listingId uint) (model.ListingResponse,error) {
	listing:=model.Listing{}
	if err := lu.lr.GetMyListingById(&listing, userId, listingId); err != nil {
		return model.ListingResponse{},err
	}
	resListing:=model.ListingResponse{
		ID: listing.ID,
		BookTitle: listing.BookTitle,
		CreatedAt: listing.CreatedAt,
		UpdatedAt: listing.UpdatedAt,
	}
	return resListing,nil
}

func (lu *listingUsecase) CreateListing(listing model.Listing)(model.ListingResponse,error){
	if err:=lu.tv.ListingValidate(listing);err!=nil{
		return model.ListingResponse{},err
	}
	if err:= lu.lr.CreateListing(&listing); err != nil {
		return model.ListingResponse{},err
	}
	resListing:=model.ListingResponse{
		ID: listing.ID,
		BookTitle: listing.BookTitle,
		CreatedAt: listing.CreatedAt,
		UpdatedAt: listing.UpdatedAt,
	}
	return resListing,nil
}

func (lu *listingUsecase) UpdateListing(listing model.Listing, userId uint, listingId uint) (model.ListingResponse,error) {
	if err:=lu.tv.ListingValidate(listing);err!=nil{
		return model.ListingResponse{},err
	}
		if err := lu.lr.UpdateListing(&listing, userId, listingId); err != nil {
		return model.ListingResponse{},err
	}
	resListing:=model.ListingResponse{
		ID: listing.ID,
		BookTitle: listing.BookTitle,
		CreatedAt: listing.CreatedAt,
		UpdatedAt: listing.UpdatedAt,
	}
	return resListing,nil
}

func (lu *listingUsecase) DeleteListing(userId uint, listingId uint) error {
	if err := lu.lr.DeleteListing(userId, listingId); err != nil {
		return err
	}
	return nil
}