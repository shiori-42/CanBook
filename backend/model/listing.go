package model

import (
	"time"
)

type Listing struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	User       User      `json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	CourseName string    `json:"course_name"`
	Instructor string    `json:"instructor"`
	BookTitle  string    `json:"book_title" gorm:"not null"`
	Condition  string    `json:"condition"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ListingResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	BookTitle string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
