/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   item_service.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori <shiori@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 16:42:06 by shiori0123        #+#    #+#             */
/*   Updated: 2024/07/13 23:05:15 by shiori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package service

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/shiori-42/textbook_change_app/go/backend/model"
	"github.com/shiori-42/textbook_change_app/go/backend/repository"
	"github.com/shiori-42/textbook_change_app/go/backend/util"
	"github.com/shiori-42/textbook_change_app/go/backend/validator"
)

func CreateItem(c echo.Context, userID uint) (model.Item, error) {
	var item model.Item
	textname := c.FormValue("text_name")
	className := c.FormValue("class_name")
	priceStr := c.FormValue("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		return item, fmt.Errorf("invalid price: %v", err)
	}
	sellType := c.FormValue("sell_type")
	image, err := c.FormFile("image")
	if err != nil {
		return item, fmt.Errorf("failed to get image file: %v", err)
	}
	src, err := image.Open()
	if err != nil {
		return item, fmt.Errorf("failed to open image file: %v", err)
	}
	defer src.Close()
	ImageName, err := util.SaveImage(src, image)
	if err != nil {
		return item, fmt.Errorf("failed to save image: %v", err)
	}
	item = model.Item{
		TextName:  textname,
		ClassName: className,
		Price:     price,
		SellType:  sellType,
		ImageName: ImageName,
		UserID:    userID,
	}
	if err := validator.ItemValidate(item); err != nil {
		return item, fmt.Errorf("validation failed: %v", err)
	}
	if err := repository.CreateItem(&item); err != nil {
		return item, fmt.Errorf("failed to create item in database: %v", err)
	}
	newItem, err := repository.GetItemByID(item.ID)
	if err != nil {
		return item, fmt.Errorf("failed to get updated item: %v", err)
	}
	return newItem, nil
}

func GetItemByID(itemID int) (model.Item, error) {
	item, err := repository.GetItemByID(itemID)
	if err != nil {
		return item, err
	}
	return item, nil
}

func UpdateItem(c echo.Context, itemID int, userID uint) (model.Item, error) {
	var item model.Item
	existingItem,err:=repository.GetItemByID(itemID)
	if err!=nil{
		return item,fmt.Errorf("failed to get existing item: %v", err)
	}
	textname := c.FormValue("text_name")
	className := c.FormValue("class_name")
	priceStr := c.FormValue("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		return item, fmt.Errorf("invalid price: %v", err)
	}
	sellType := c.FormValue("sell_type")
	var ImageName string
	image, err := c.FormFile("image")
	if err == nil {
		src, err := image.Open()
		if err != nil {
			return item, fmt.Errorf("failed to open image file: %v", err)
		}
		defer src.Close()
		ImageName, err = util.SaveImage(src, image)
		if err != nil {
			return item, fmt.Errorf("failed to save image: %v", err)
		}
		existingItem.ImageName = ImageName
	}
	existingItem.ImageName=ImageName
	existingItem.TextName=textname
	existingItem.ClassName=className
	existingItem.Price=price
	existingItem.SellType=sellType
	
	if err := validator.ItemValidate(existingItem); err != nil {
		return item, fmt.Errorf("validation failed: %v", err)
	}
	if err := repository.UpdateItem(&existingItem, itemID, userID); err != nil {
		return item, fmt.Errorf("failed to update item in database: %v", err)
	}
	updatedItem, err := repository.GetItemByID(itemID)
	if err != nil {
		return item, fmt.Errorf("failed to get updated item: %v", err)
	}
	return updatedItem, nil
}

func DeleteItem(itemID string, userID uint) error {
	if err := repository.DeleteItem(itemID, userID); err != nil {
		return err
	}
	return nil
}

func SearchItemsByKeyword(keyword string) (model.Items, error) {
	items, err := repository.SearchItemsByKeyword(keyword)
	if err != nil {
		return items, err
	}
	return items, nil
}

func SearchItemsByCollege(college string) (model.Items, error) {
	items, err := repository.SearchItemsByCollege(college)
	if err != nil {
		return items, err
	}
	return items, nil
}
