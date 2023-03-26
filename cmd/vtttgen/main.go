package main

import (
	"log"
	"os"

	"github.com/umemak/vtttgen"
	"github.com/urfave/cli/v2"
)

const version = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "vtttgen"
	app.Usage = "WEBVTT Thumbnail GENerator"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "target",
			Value: "",
			Usage: "target dir name",
		},
		&cli.Int64Flag{
			Name:  "columns",
			Value: 6,
			Usage: "thumbnail columns",
		},
		&cli.Int64Flag{
			Name:  "rows",
			Value: 10,
			Usage: "thumbnail rows",
		},
	}
	app.Action = vtttgen.Run
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
