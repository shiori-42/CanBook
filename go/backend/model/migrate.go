/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   migrate.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori <shiori@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 12:02:59 by shiori0123        #+#    #+#             */
/*   Updated: 2024/07/13 23:43:12 by shiori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package model

import (
	"log"
)

func AutoMigrate() error {
	err := AutoMigrateUser()
	if err != nil {
		log.Printf("Error migrating user table: %v", err)
		return err
	}

	err = AutoMigrateItem()
	if err != nil {
		log.Printf("Error migrating item table: %v", err)
		return err
	}
	
	// err = AutoMigrateLikes()
	// if err != nil {
	// 	log.Printf("Error migrating  likes table: %v", err)
	// 	return err
	// }

	return nil
}
