//go:build tools

package task

// Package tools tracks dependencies for tools that used in the build process.
// See https://github.com/golang/go/wiki/Modules
import (
	_ "golang.org/x/tools/cmd/goimports"
)
