package main

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"

	"github.com/urfave/cli"
)

//go:generate swagger generate spec -m -i ../../swagger-basic.yml -o ../../swagger.json
//go:generate swagger flatten ../../swagger.json -o ../../swagger.json
//go:generate swagger validate ../../swagger.json

var version string

func main() {
	app := cli.NewApp()
	app.Name = "events-api"
	app.Version = version
	app.Usage = "Events api for aggregating resources changes and events"
	app.Flags = flags

	figure.NewFigure(app.Name, "banner3", true).Print()

	app.Action = initServer

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
