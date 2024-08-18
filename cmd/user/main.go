package main

import (
	"log"
	"os"

	_ "embed"

	"github.com/urfave/cli/v2"
)

//go:generate ../../scripts/versiongen.sh ./VERSION
//go:embed VERSION
var version string

func main() {
	app := cli.App{
		Name:    "Account Manager Service",
		Usage:   "Provides services for user-related operations",
		Version: version,
		Action:  action,
		Flags:   flags,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
