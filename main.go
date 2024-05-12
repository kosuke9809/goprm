package main

import (
	"log"
	"os"

	"github.com/kosuke9809/goprm/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	goprm := &cli.App{
		Name:  "goprm",
		Usage: "Go Simple Parameter Manager",
		Commands: []*cli.Command{
			cmd.InitCommand,
		},
	}
	err := goprm.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
