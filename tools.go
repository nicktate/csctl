// +build tools

// Package tools is used to import go modules that we use for tooling as dependencies.
// For more information, please refer to: https://github.com/go-modules-by-example/index/blob/ac9bf72/010_tools/README.md
package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"

	_ "github.com/go-openapi/runtime"
	_ "github.com/jessevdk/go-flags"
)
