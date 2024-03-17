package model

import (
	"time"
)

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	UserName   string    `json:"username"`
	Email      string    `json:"email" gorm:"unique"`
	SchoolName string    `json:"schoolname"`
	Campus     string    `json:"campus"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserName string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
}
