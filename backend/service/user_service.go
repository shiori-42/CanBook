/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   user_service.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 18:49:38 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/20 21:17:27 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package service

import (
	"fmt"
	"github.com/shiori-42/textbook_change_app/model"
	"github.com/shiori-42/textbook_change_app/repository"
	"github.com/shiori-42/textbook_change_app/validator"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(user model.User) (model.UserResponse, error) {
	var userRes model.UserResponse
	if err := validator.UserValidate(user); err != nil {
		return userRes, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return userRes, err
	}
	user.Password = string(hash)

	if err := repository.CreateUser(&user); err != nil {
		return userRes, err
	}
	resUser := model.UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}
	return resUser, nil
}

func LoginUser(user model.User) (model.User, error) {
    if err := validator.UserValidate(user); err != nil {
        return model.User{}, err
    }

    storedUser, err := repository.GetUserByEmail(user.Email)
    if err != nil {
        return model.User{}, err
    }
    if storedUser == nil {
        return model.User{}, fmt.Errorf("user not found")
    }

    err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
    if err != nil {
        return model.User{}, err
    }

    return *storedUser, nil
}