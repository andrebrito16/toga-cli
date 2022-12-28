package gpg

import (
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "gpg",
		Usage: "run gpg to encrypt and decrypt files",
		Subcommands: []*cli.Command{
			{
				Name:     "encrypt",
				Usage:    "encrypt a file with gpg",
				Category: "encryption",
				Action:   EncryptSymetrical,
				Aliases:  []string{"e"},
			},
		},
	}
}
