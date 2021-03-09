// Copyright 2021 Harran Ali. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"gorm.io/gorm"
)

// Product is model that reperesents a database table
type Product struct {
	gorm.Model
	Code  string
	Price int
}
