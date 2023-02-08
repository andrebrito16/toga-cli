package git

import (
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "git",
		Usage: "run usefull git commands",
		Subcommands: []*cli.Command{
			{
				Name:     "gone",
				Aliases:  []string{"g"},
				Usage:    "remove all local branches that not exist on remote",
				Category: "git",
				Action:   GitGone,
			},
		},
	}
}
