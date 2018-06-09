package main

import (
	"os"
	"path/filepath"

	"github.com/kunickiaj/vault-issuer/pkg/cmd"
)

func main() {
	baseName := filepath.Base(os.Args[0])
	c := cmd.NewCommand(baseName)

	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
