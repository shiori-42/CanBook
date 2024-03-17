package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/shiori-42/textbook_change_app/model"
)

type IListingRepository interface {
	GetAllMyListings(listigs *[]model.Listing, userId uint) error
	GetMyListingById(listing *model.Listing, userId uint, listingId uint) error
	CreateListing(listing *model.Listing) error
	UpdateListing(listing *model.Listing, userId uint, listingId uint) error
	DeleteListing(userId uint, listingId uint) error
}

type listingRepository struct {
	db *gorm.DB
}

func NewListingRepository(db *gorm.DB) IListingRepository {
	return &listingRepository{db}
}

func (lr *listingRepository) GetAllMyListings(listings *[]model.Listing, userId uint) error {
	if err := lr.db.Joins("User").Where("user_id = ?", userId).Order("created_at").Find(listings).Error; err != nil {
		return err
	}
	return nil
}

func (lr *listingRepository) GetMyListingById(listing *model.Listing, userId uint, listingId uint) error {
	if err := lr.db.Joins("User").Where("user_id = ? AND id = ?", userId, listingId).First(listing).Error; err != nil {
		return err
	}
	return nil
}

func (lr *listingRepository) CreateListing(listing *model.Listing) error {
	if err := lr.db.Create(listing).Error; err != nil {
		return err
	}
	return nil
}

func (lr *listingRepository) UpdateListing(listing *model.Listing, userId uint, listingId uint) error {
	result := lr.db.Model(listing).Clauses(clause.Returning{}).Where("id=? AND user_id=?", listingId, userId).Update("book_title", listing.BookTitle)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (lr *listingRepository) DeleteListing(userId uint, listingId uint) error {
	result := lr.db.Where("id=? AND user_id=?", listingId, userId).Delete(&model.Listing{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
