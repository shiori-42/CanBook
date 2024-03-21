/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   item_repository.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 16:42:38 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/20 21:56:57 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package repository

import (
	"fmt"
	"database/sql"
	"github.com/shiori-42/textbook_change_app/db"
	"github.com/shiori-42/textbook_change_app/model"
)

func GetAllItems() (model.Items, error) {
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
	`

	rows, err := db.DB.Query(query)
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

func GetItemByID(itemID string) (model.Item, error) {
	var item model.Item
	query := `
		SELECT
			items.id,
			items.name,
			items.category_id,
			categories.name AS category,
			items.image_name,
			items.user_id
		FROM items
		JOIN categories ON items.category_id = categories.id
		WHERE items.id = ?
	`

	err := db.DB.QueryRow(query, itemID).Scan(
        &item.ID,
        &item.Name,
        &item.CategoryID,
        &item.Category,
        &item.ImageName,
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
        INSERT INTO items (name, category_id, image_name)
        VALUES (?, ?, ?)
    `
    _, err := db.DB.Exec(query, item.Name, item.CategoryID, item.ImageName)
    if err != nil {
        return err
    }
    return nil
}

func UpdateItem(item model.Item) error {
	query := `
		UPDATE items
		SET name = ?, category_id = ?, image_name = ?
		WHERE id = ?
	`
	_, err := db.DB.Exec(query, item.Name, item.CategoryID, item.ImageName, item.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteItem(itemID string, userID uint) error {
    query := `DELETE FROM items WHERE id = ? AND user_id = ?`
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
		WHERE items.name LIKE ?
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

func CheckCategoryIDExist(categoryID int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM categories WHERE id = ?)"
	err := db.DB.QueryRow(query, categoryID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}