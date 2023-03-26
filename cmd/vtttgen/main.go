package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const version = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "vtttgen"
	app.Usage = "WEBVTT Thumbnail GENerator"
	app.Flags = []cli.Flag{
		&cli.Int64Flag{
			Name:  "width",
			Value: 240,
			Usage: "thumbnail width",
		},
		&cli.Int64Flag{
			Name:  "height",
			Value: 120,
			Usage: "thumbnail height",
		},
		&cli.Int64Flag{
			Name:  "columns",
			Value: 10,
			Usage: "thumbnail columns",
		},
		&cli.Int64Flag{
			Name:  "rows",
			Value: 6,
			Usage: "thumbnail rows",
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
