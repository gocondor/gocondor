// Copyright 2021 Harran Ali. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package database

import (
	"fmt"
	"log"
	"os"

	"github.com/gincoat/gincoat/core/env"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// New initiates database connection
func New() *gorm.DB {
	//initiate database drier
	switch env.Get("DB_DRIVER") {
	case "mysql":
		db, err := prepareMysql()
		if err != nil {
			log.Fatal(err)
		}
		return db
	case "sqlite":
		dbName := env.Get("SQLITE_FILE")
		_, err := os.Stat("../../" + dbName)
		if err != nil {
			panic(err)
		}
		db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		return db
	default:
		db, err := prepareMysql()
		if err != nil {
			log.Fatal(err)
		}
		return db
	}
}

func prepareMysql() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		env.Get("MYSQL_USERNAME"),
		env.Get("MYSQL_PASSWORD"),
		env.Get("MYSQL_HOST"),
		env.Get("MYSQL_PORT"),
		env.Get("MYSQL_DB_NAME"),
		env.Get("MYSQL_CHARSET"),
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Resolve() *gorm.DB {
	return db
}
