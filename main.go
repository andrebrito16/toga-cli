package main

import (
	"os"

	"github.com/andrebrito16/toga-cli/git"
	"github.com/andrebrito16/toga-cli/gpg"
	"github.com/andrebrito16/toga-cli/terminal"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name: "toga",
		Authors: []*cli.Author{
			{
				Name:  "André Brito",
				Email: "felipe.brito016@gmail.com",
			},
		},
		Usage: "Just a simple CLI for setup dev environment. Don't lose your time, just takeoff!",
		Commands: []*cli.Command{
			terminal.Command(),
			gpg.Command(),
			git.Command(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.WithError(err).Error("Failed to run CLI")
		os.Exit(1)
	}
}
