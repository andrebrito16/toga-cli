package git

import (
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "git",
		Usage: "git utils",
		Subcommands: []*cli.Command{
			{
				Name:     "gone",
				Usage:    "encrypt a file with gpg",
				Category: "encryption",
				Action:   CleanBranchLocal,
				Aliases:  []string{"g"},
			},
		},
	}
}
