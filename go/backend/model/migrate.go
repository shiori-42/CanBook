/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   migrate.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 12:02:59 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/26 21:10:38 by shiori0123       ###   ########.fr       */
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

	return nil
}
