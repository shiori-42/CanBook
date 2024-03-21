/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   item.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 12:02:51 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/20 16:57:42 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package model

import (
	"time"

	"github.com/shiori-42/textbook_change_app/db"
)

type Item struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CategoryID int       `json:"category_id"`
	Category   string    `json:"category"`
	ImageName  string    `json:"imagename"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	UserID     uint      `json:"user_id"`
}

type Items struct {
	Items []Item `json:"items"`
}

type ItemResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func AutoMigrateItem() error {
	_, err := db.DB.Exec(`
		CREATE TABLE IF NOT EXISTS items (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name TEXT NOT NULL,
			category_id INT NOT NULL,
			image_name TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			user_id INT,
			FOREIGN KEY (category_id) REFERENCES categories(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`)
	if err != nil {
		return err
	}
	return nil
}
