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
		Name:    "Payment Manager Service",
		Usage:   "Provides services for payment-related operations",
		Version: version,
		Action:  action,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
