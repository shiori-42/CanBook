/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   user_repository.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/29 05:11:19 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/29 05:13:51 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package repository

import (
	"database/sql"

	"github.com/shiori-42/textbook_change_app/go/backend/db"
	"github.com/shiori-42/textbook_change_app/go/backend/model"
)

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	query := `SELECT id, name, email, password, college, campus, created_at, updated_at FROM users WHERE email = $1`
	err := db.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.College, &user.Campus, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *model.User) error {
	query := `INSERT INTO users (name, email, password, college, campus) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.DB.QueryRow(query, user.Name, user.Email, user.Password, user.College, user.Campus).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByID(userID uint) (model.User, error) {
	var user model.User
	query := `SELECT * FROM users WHERE id = $1`
	err := db.DB.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.Email, &user.Password, user.College, user.Campus, &user.CreatedAt, &user.UpdatedAt)
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
