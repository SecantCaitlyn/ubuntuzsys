//go:build tools
// +build tools

package main

import (
	"os"

	"github.com/ubuntu/zsys/internal/generators"
)

func main() {
	if generators.InstallOnlyMode() {
		os.Exit(1)
	}
}
