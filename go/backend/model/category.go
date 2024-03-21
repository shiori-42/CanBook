/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   category.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 12:02:22 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/20 12:02:33 by shiori0123       ###   ########.fr       */
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
			id INT AUTO_INCREMENT PRIMARY KEY,
			name TEXT NOT NULL
		)
	`)
	if err != nil {
		return err
	}
	return nil
}
