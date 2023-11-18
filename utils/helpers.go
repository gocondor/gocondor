// Copyright 2023 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package utils

import (
	"crypto/md5"
	"fmt"
)

// generate a hashed string to be used as key for caching auth jwt token
func CreateAuthTokenHashedCacheKey(userID uint, userAgent string) string {
	cacheKey := fmt.Sprintf("userid:_%v_useragent:_%v_jwt_token", userID, userAgent)
	hashedCacheKey := fmt.Sprintf("%v", fmt.Sprintf("%x", md5.Sum([]byte(cacheKey))))

	return hashedCacheKey
}
