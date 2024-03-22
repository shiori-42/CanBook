/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   user_repository.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 18:47:02 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/22 06:58:17 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package repository

import (
	"database/sql"
	"github.com/shiori-42/textbook_change_app/db"
	"github.com/shiori-42/textbook_change_app/model"
)

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	query := `SELECT * FROM users WHERE email = ?`
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
	query := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`
	result, err := db.DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = uint(id)
	return nil
}

func GetUserByID(userID uint) (model.User, error) {
	var user model.User
	query := `SELECT * FROM users WHERE id = ?`
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
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)"
	err := db.DB.QueryRow(query, userID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
