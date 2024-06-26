// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/module_include_list/module_include_list.go - DO NOT EDIT.

package include

import (
	// Import packages to perform 'func InitializeModule()' when in-use.
	m0 "github.com/elastic/beats/v7/x-pack/auditbeat/processors/sessionmd"

	// Import packages that perform 'func init()'.
	_ "github.com/elastic/beats/v7/x-pack/auditbeat/module/system"
	_ "github.com/elastic/beats/v7/x-pack/auditbeat/module/system/host"
	_ "github.com/elastic/beats/v7/x-pack/auditbeat/module/system/login"
	_ "github.com/elastic/beats/v7/x-pack/auditbeat/module/system/package"
	_ "github.com/elastic/beats/v7/x-pack/auditbeat/module/system/process"
	_ "github.com/elastic/beats/v7/x-pack/auditbeat/module/system/socket"
	_ "github.com/elastic/beats/v7/x-pack/auditbeat/module/system/user"
)

// InitializeModules initialize all of the modules.
func InitializeModule() {
	m0.InitializeModule()
}
