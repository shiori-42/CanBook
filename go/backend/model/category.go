/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   category.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 12:02:22 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/22 11:37:48 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package model

import (
	"github.com/shiori-42/textbook_change_app/db"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func AutoMigrateCategory() error {
	_, err := db.DB.Exec(`
        CREATE TABLE IF NOT EXISTS categories (
            id SERIAL PRIMARY KEY,
            name VARCHAR(50) NOT NULL,
            CONSTRAINT categories_name_key UNIQUE (name)
        )
    `)
	if err != nil {
		return err
	}

	categories := []string{"Rental", "Sell", "Buy"}
	for _, category := range categories {
		var categoryId int
		insertSQL := `INSERT INTO categories (name) VALUES ($1) RETURNING id`
		err := db.DB.QueryRow(insertSQL, category).Scan(&categoryId)
		if err != nil {
			if err.Error() == `pq: duplicate key value violates unique constraint "categories_name_key"` {
				continue
			}
			return err
		}
	}
	return nil
}
