package repository

import (
	"database/sql"
	"github.com/shiori-42/textbook_change_app/go/backend/db"
	"github.com/shiori-42/textbook_change_app/go/backend/model"
)

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	query := `SELECT * FROM users WHERE email = $1`
	err := db.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *model.User) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := db.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByID(userID uint) (model.User, error) {
	var user model.User
	query := `SELECT * FROM users WHERE id = $1`
	err := db.DB.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}
	return user, nil
}

func CheckUserExist(userID uint) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)"
	err := db.DB.QueryRow(query, userID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
