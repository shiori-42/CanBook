/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   user.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 12:03:06 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/26 21:44:34 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package model

import (
	"time"

	"github.com/shiori-42/textbook_change_app/go/backend/db"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func AutoMigrateUser() error {
	_, err := db.DB.Exec(`
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )
`)
	if err != nil {
		return err
	}
	return nil
}
