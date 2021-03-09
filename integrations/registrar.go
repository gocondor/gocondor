// Copyright 2021 Harran Ali. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package integrations

import "github.com/harranali/gincoat/core/pkgintegrator"

// RegisterPKGIntegrations registers the integrations to gin.Context
func RegisterPKGIntegrations() {
	pkgi := pkgintegrator.Resolve()

	//add your packges here
	pkgi.Integrate(PKGIntegratorExample)

}
