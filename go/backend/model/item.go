/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   item.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori <shiori@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 12:02:51 by shiori0123        #+#    #+#             */
/*   Updated: 2024/07/13 20:30:10 by shiori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package model

import (
	"time"

	"github.com/shiori-42/textbook_change_app/go/backend/db"
)

type Item struct {
	ID        int       `json:"id"`
	TextName  string    `json:"text_name"`
	ClassName string    `json:"class_name"`
	Price     int       `json:"price"`
	SellType  string    `json:"sell_type"`
	ImageName string    `json:"image_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint      `json:"user_id"`
}

type Items struct {
	Items []Item `json:"items"`
}

type ItemResponse struct {
	ID        int       `json:"id"`
	TextName  string    `json:"text_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func AutoMigrateItem() error {
	_, err := db.DB.Exec(`
    CREATE TABLE IF NOT EXISTS items (
        id SERIAL PRIMARY KEY,
        text_name TEXT NOT NULL,
        class_name TEXT NOT NULL,
        image_name TEXT NOT NULL,
        price INT NOT NULL,
        sell_type TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        user_id INT NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id)
    )
`)
	if err != nil {
		return err
	}
	return nil
}
