// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// New initiates database connection
func New() *gorm.DB {
	//initiate database drier
	switch os.Getenv("DB_DRIVER") {
	case "mysql":
		db, err := prepareMysql()
		if err != nil {
			log.Fatal(err)
		}
		return db
	case "sqlite":
		dbName := os.Getenv("SQLITE_FILE")
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
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB_NAME"),
		os.Getenv("MYSQL_CHARSET"),
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
