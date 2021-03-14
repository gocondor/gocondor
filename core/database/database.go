// Copyright 2021 Harran Ali. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package database

import (
	"fmt"
	"log"

	"github.com/harranali/gincoat/core/env"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
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
	case "postgresql":
		db, err := preparePostgresql()
		if err != nil {
			log.Fatal(err)
		}
		return db
	default:
		db, err := prepareMysql()
		if err != nil {
			log.Fatal(err)
		}
		return db
	}

	return db
}

func preparePostgresql() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		env.Get("POSTGRESQL_HOST"),
		env.Get("POSTGRESQL_USERNAME"),
		env.Get("POSTGRESQL_PASSWORD"),
		env.Get("POSTGRESQL_DB_NAME"),
		env.Get("POSTGRESQL_PORT"),
		env.Get("POSTGRESQL_SSL_MODE"),
		env.Get("POSTGRESQL_TIMEZONE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
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
	//create new db
	return db
}
