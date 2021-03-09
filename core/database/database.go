// Copyright 2021 Harran Ali. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package database

import (
	"fmt"
	"log"

	"github.com/harranali/gincoat/core/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// New initiates database connection
func New() *gorm.DB {
	//initiate database drier
	switch env.Get("DB_DRIVER") {
	case "mysql":
		//example "root:@tcp(127.0.0.1:3306)/gincoat?charset=utf8mb4&parseTime=True&loc=Local"
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/gincoat?charset=%s&parseTime=True&loc=Local",
			env.Get("MYSQL_USERNAME"),
			env.Get("MYSQL_PASSWORD"),
			env.Get("MYSQL_HOST"),
			env.Get("MYSQL_PORT"),
			env.Get("MYSQL_CHARSET"),
		)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal(err)
		}

		return db
	}

	return db
}

func Resolve() *gorm.DB {
	//create new db
	return db
}
