/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   db.go                                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 12:02:14 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/22 08:43:35 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func InitDB(dataSourceName string) error {
	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	return DB.Ping()
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalln(err)
	}
}
