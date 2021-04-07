// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package integrations

import "github.com/gocondor/core/pkgintegrator"

// RegisterPKGIntegrations registers the integrations
func RegisterPKGIntegrations() {
	pkgi := pkgintegrator.Resolve()

	//add your packges here
	pkgi.Integrate(PKGIntegratorExample)

}
