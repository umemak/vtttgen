package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

const version = "0.0.0"

func main() {
	app := cli.NewApp()
	app.Name = "vtttgen"
	app.Usage = "WEBVTT Thumbnail GENerator"
	app.Flags = []cli.Flag{
		cli.Int64Flag{
			Name:  "width, w",
			Value: 240,
			Usage: "thumbnail width",
		},
		cli.Int64Flag{
			Name:  "height, h",
			Value: 120,
			Usage: "thumbnail height",
		},
		cli.Int64Flag{
			Name:  "columns, c",
			Value: 10,
			Usage: "thumbnail columns",
		},
		cli.Int64Flag{
			Name:  "rows, r",
			Value: 6,
			Usage: "thumbnail rows",
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
