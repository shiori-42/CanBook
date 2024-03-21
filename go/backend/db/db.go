/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   db.go                                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 12:02:14 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/20 22:30:25 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	config := mysql.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")),
		DBName:               os.Getenv("MYSQL_DATABASE"),
		AllowNativePasswords: true,
	}

	DB, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalln(err)
	}
}
