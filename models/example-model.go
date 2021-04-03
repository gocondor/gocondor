// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"gorm.io/gorm"
)

// User is an example model
type User struct {
	gorm.Model
	name     string
	email    string
	password string
	age      int
}
