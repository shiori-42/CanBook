/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   item_service.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 16:42:06 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/20 21:58:35 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/shiori-42/textbook_change_app/model"
	"github.com/shiori-42/textbook_change_app/repository"
	"github.com/shiori-42/textbook_change_app/util"
	"github.com/shiori-42/textbook_change_app/validator"
)

func CreateItem(c echo.Context) (model.Item, error) {
	var item model.Item
	if err := c.Bind(&item); err != nil {
		return item, err
	}
	if err := validator.ItemValidate(item); err != nil {	
		return item, err	
		}	
		
	categoryExist, err := repository.CheckCategoryIDExist(item.CategoryID)
	if err != nil {
		return item, err
	}
	if !categoryExist {
		return item, fmt.Errorf("category_id %d does not exist", item.CategoryID)
	}
	file, err := c.FormFile("image")
	if err != nil {
		return item, err
	}
	src, err := file.Open()
	if err != nil {
		return item, err
	}
	defer src.Close()
	item.ImageName, err = util.SaveImage(src)
	if err != nil {
		return item, err
	}
	return item, nil
}

func GetItemByID(itemID string) (model.Item, error) {
	item, err := repository.GetItemByID(itemID)
	if err != nil {
		return item, err
	}
	return item, nil
}

func UpdateItem(c echo.Context) (model.Item, error) {
	itemID:=c.Param("itemId")
	var item model.Item
	if err := c.Bind(&item); err != nil {
		return item, err
	}
	if err := validator.ItemValidate(item); err != nil {	
		return item, err	
		}
	_, err := repository.GetItemByID(itemID)
	if err != nil {
		return item, err
	}
	if item.CategoryID != 0 {
		categoryExist, err := repository.CheckCategoryIDExist(item.CategoryID)
		if err != nil {
			return item, err
		}
		if !categoryExist {
			return item, fmt.Errorf("category_id %d does not exist", item.CategoryID)
		}
	}

	file, err := c.FormFile("image")
	if err == nil {
		src, err := file.Open()
		if err != nil {
			return item, err
		}
		defer src.Close()

		item.ImageName, err = util.SaveImage(src)
		if err != nil {
			return item, err
		}
	}

	if err := repository.UpdateItem(item); err != nil {
		return item, err
	}
	updatedItem, err := repository.GetItemByID(itemID)
    if err != nil {
        return item, err
    }
    return updatedItem, nil
}

func DeleteItem(itemID string, userID uint) error {
	if err := repository.DeleteItem(itemID, userID); err != nil {
		return err
	}
	return nil
}

func SearchItems(keyword string) (model.Items, error) {
	items, err := repository.SearchItemsByKeyword(keyword)
	if err != nil {
		return items, err
	}
	return items, nil
}