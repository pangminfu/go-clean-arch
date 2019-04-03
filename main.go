package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var cmds []cli.Command

func main() {
	app := cli.NewApp()
	app.Name = "GO Clean Architecture Demo CLI App"
	app.Usage = "Demo CLI tool"
	app.Version = "1.0.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name: "pangminfu",
		},
	}
	app.Commands = cmds

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	return
}
