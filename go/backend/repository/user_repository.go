/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   user_repository.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 18:47:02 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/20 21:15:52 by shiori0123       ###   ########.fr       */
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
    err := db.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
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