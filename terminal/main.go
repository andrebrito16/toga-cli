package terminal

import (
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "terminal",
		Usage: "run commands to configure your terminal",
		Subcommands: []*cli.Command{
			{
				Name:  "zsh",
				Usage: "configure zsh",
				Subcommands: []*cli.Command{
					{
						Name:   "install",
						Usage:  "Install zsh on your terminal",
						Action: ZshInstall,
					},
				},
			},
			{
				Name:  "fish",
				Usage: "configure fish terminal",
				Subcommands: []*cli.Command{
					{
						Name:   "install",
						Usage:  "Install fish on your terminal",
						Action: FishInstall,
					},
				},
			},
		},
	}
}
