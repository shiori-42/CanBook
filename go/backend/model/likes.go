/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   likes.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori <shiori@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/07/13 19:53:37 by shiori            #+#    #+#             */
/*   Updated: 2024/07/13 20:39:46 by shiori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package model

import (
	"github.com/shiori-42/textbook_change_app/go/backend/db"
)

type Like struct {
    UserID int `json:"user_id"`
    ItemID int `json:"item_id"`
}

func AutoMigrateLikes() error {
	_, err := db.DB.Exec(`
    CREATE TABLE likes (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    item_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, item_id)
	)
`)
	if err != nil {
		return err
	}
	return nil
}
