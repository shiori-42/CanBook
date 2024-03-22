/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   item_repository.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 16:42:38 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/22 15:52:47 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package repository

import (
	"database/sql"
	"fmt"
	"github.com/shiori-42/textbook_change_app/go/backend/db"
	"github.com/shiori-42/textbook_change_app/go/backend/model"
)

func GetAllItems() (model.Items, error) {
	var items model.Items
	query := `
        SELECT 
            items.id, 
            items.name, 
            items.category_id, 
            categories.name AS category, 
            items.image_name,
            items.created_at,
            items.updated_at,
            items.user_id
        FROM items
        JOIN categories ON items.category_id = categories.id
    `

	rows, err := db.DB.Query(query)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.Item
		err = rows.Scan(&item.ID, &item.Name, &item.CategoryID, &item.Category, &item.ImageName, &item.CreatedAt, &item.UpdatedAt, &item.UserID)
		if err != nil {
			return items, err
		}
		items.Items = append(items.Items, item)
	}

	return items, nil
}

func GetItemByID(itemID int) (model.Item, error) {
	var item model.Item
	query := `
        SELECT items.id, items.name, items.category_id, categories.name AS category_name, items.image_name, items.created_at, items.updated_at, items.user_id
        FROM items
        JOIN categories ON items.category_id = categories.id
        WHERE items.id = $1
    `

	err := db.DB.QueryRow(query, itemID).Scan(
		&item.ID,
		&item.Name,
		&item.CategoryID,
		&item.Category,
		&item.ImageName,
		&item.CreatedAt,
		&item.UpdatedAt,
		&item.UserID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, nil
		}
		return item, err
	}

	return item, nil
}

func CreateItem(item *model.Item) error {
	query := `
        INSERT INTO items (name, category_id, image_name, user_id)
        VALUES ($1, $2, $3, $4) RETURNING id
    `
	err := db.DB.QueryRow(query, item.Name, item.CategoryID, item.ImageName, item.UserID).Scan(&item.ID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateItem(item *model.Item, itemID int, userID uint) error {
	query := `
        UPDATE items
        SET name = $1, category_id = $2, image_name = $3
        WHERE id = $4 AND user_id = $5
    `
	_, err := db.DB.Exec(query, item.Name, item.CategoryID, item.ImageName, itemID, userID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteItem(itemID string, userID uint) error {
	query := `DELETE FROM items WHERE id = $1 AND user_id = $2`
	result, err := db.DB.Exec(query, itemID, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("item not found or not owned by the user")
	}
	return nil
}

func SearchItemsByKeyword(keyword string) (model.Items, error) {
	var items model.Items
	query := `
        SELECT
            items.id,
            items.name,
            items.category_id,
            categories.name AS category,
            items.image_name
        FROM items
        JOIN categories ON items.category_id = categories.id
        WHERE items.name LIKE $1
    `
	rows, err := db.DB.Query(query, "%"+keyword+"%")
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.Item
		err = rows.Scan(&item.ID, &item.Name, &item.CategoryID, &item.Category, &item.ImageName)
		if err != nil {
			return items, err
		}
		items.Items = append(items.Items, item)
	}

	return items, nil
}
