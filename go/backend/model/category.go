/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   category.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 12:02:22 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/21 18:05:00 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package model

import (
	"database/sql"
	"github.com/shiori-42/textbook_change_app/db"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func AutoMigrateCategory() error {
	_, err := db.DB.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(50) NOT NULL
		)
	`)
	if err != nil {
		return err
	}
	categories := []string{"Rental", "Sell", "Buy"}
	for _, category := range categories {
		var categoryId int
		err := db.DB.QueryRow("SELECT id FROM categories WHERE name = ?", category).Scan(&categoryId)
		if err == sql.ErrNoRows {
			_, err = db.DB.Exec("INSERT INTO categories (name) VALUES (?)", category)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}
	return nil
}
