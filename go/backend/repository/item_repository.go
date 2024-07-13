/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   item_repository.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori <shiori@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 16:42:38 by shiori0123        #+#    #+#             */
/*   Updated: 2024/07/13 22:51:00 by shiori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package repository

import (
	"database/sql"
	"fmt"

	"github.com/shiori-42/textbook_change_app/go/backend/db"
	"github.com/shiori-42/textbook_change_app/go/backend/model"
)

func GetMyItems(userID uint) (model.Items, error) {
	var items model.Items
	query := `
        SELECT 
            items.id, 
            items.text_name,
			items.class_name, 
            items.price,
            items.sell_type,
            items.image_name,
            items.created_at,
            items.updated_at,
            items.user_id
        FROM items
        WHERE items.user_id = $1
    `

	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.Item
		err = rows.Scan(
			&item.ID,
			&item.TextName,
			&item.ClassName,
			&item.Price,
			&item.SellType,
			&item.ImageName,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.UserID,
		)
		if err != nil {
			return items, err
		}
		items.Items = append(items.Items, item)
	}

	return items, nil
}

func GetAllUserItems() (model.Items, error) {
	var items model.Items
	query := `
        SELECT 
            items.id, 
            items.text_name,
			items.class_name, 
            items.price,
            items.sell_type,
            items.image_name,
            items.created_at,
            items.updated_at,
            items.user_id
        FROM items
    `

	rows, err := db.DB.Query(query)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.Item
		err = rows.Scan(
			&item.ID,
			&item.TextName,
			&item.ClassName,
			&item.Price,
			&item.SellType,
			&item.ImageName,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.UserID,
		)
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
        SELECT 
            items.id, 
            items.text_name,
			items.class_name, 
            items.price,
            items.sell_type,
            items.image_name, 
            items.created_at, 
            items.updated_at, 
            items.user_id
        FROM items
        WHERE items.id = $1
    `

	err := db.DB.QueryRow(query, itemID).Scan(
		&item.ID,
		&item.TextName,
		&item.ClassName,
		&item.Price,
		&item.SellType,
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
        INSERT INTO items (text_name, class_name , image_name, price, sell_type, user_id)
        VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
    `
	err := db.DB.QueryRow(query, item.TextName, item.ClassName, item.ImageName, item.Price, item.SellType, item.UserID).Scan(&item.ID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateItem(item *model.Item, itemID int, userID uint) error {
	query := `
        UPDATE items
        SET text_name = $1,  class_name= $2, image_name = $3, price = $4, sell_type = $5
        WHERE id = $6 AND user_id = $7
    `
	_, err := db.DB.Exec(query, item.TextName, item.ClassName, item.ImageName, item.Price, item.SellType, itemID, userID)
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
            i.id, 
            i.text_name,
			i.class_name,

            i.price,
            i.sell_type,
            i.image_name,
            i.created_at,
            i.updated_at,
            i.user_id
        FROM items i
        WHERE i.text_name LIKE $1
    `
	rows, err := db.DB.Query(query, "%"+keyword+"%")
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.Item
		err = rows.Scan(
			&item.ID,
			&item.TextName,
			&item.ClassName,
			&item.Price,
			&item.SellType,
			&item.ImageName,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.UserID,
		)
		if err != nil {
			return items, err
		}
		items.Items = append(items.Items, item)
	}

	return items, nil
}

func SearchItemsByCollege(collegeName string) (model.Items, error) {
	var items model.Items
	query := `
        SELECT 
            i.id, 
            i.text_name,
            i.class_name, 
            i.price,
            i.sell_type,
            i.image_name,
            i.created_at,
            i.updated_at,
            i.user_id
        FROM items i
        JOIN users u ON i.user_id = u.id
        WHERE u.college = $1
    `

	rows, err := db.DB.Query(query, collegeName)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.Item
		err = rows.Scan(
			&item.ID,
			&item.TextName,
			&item.ClassName,
			&item.Price,
			&item.SellType,
			&item.ImageName,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.UserID,
		)
		if err != nil {
			return items, err
		}
		items.Items = append(items.Items, item)
	}

	return items, nil
}
