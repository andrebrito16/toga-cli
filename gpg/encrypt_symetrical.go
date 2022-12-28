package gpg

import (
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/andrebrito16/toga-cli/evaluator"
	"github.com/urfave/cli/v2"
)

func EncryptSymetrical(c *cli.Context) error {
	password := ""
	prompt := &survey.Password{
		Message: "Enter password:",
	}

	sourceFilePath := c.Args().Get(0)

	// Check if the file exists
	if _, err := os.Stat(sourceFilePath); os.IsNotExist(err) {
		return cli.Exit("File not found!", 1)
	}
	newFileName := sourceFilePath[:strings.LastIndex(sourceFilePath, ".")] + ".gpg"

	survey.AskOne(prompt, &password, survey.WithValidator(survey.Required))

	// Check if the file already exists
	if _, err := os.Stat(newFileName); !os.IsNotExist(err) {
		prompt := &survey.Confirm{
			Message: "File already exists. Do you want to overwrite it?",
		}

		overwrite := false

		survey.AskOne(prompt, &overwrite)

		if !overwrite {
			return cli.Exit("Aborted!", 1)
		}

		// Remove the file
		err := os.Remove(newFileName)
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}
	}

	err := evaluator.Run("gpg", []string{"--symmetric", "--cipher-algo", "AES256", "--batch", "--passphrase", password, "--output", newFileName, sourceFilePath})
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}
	return nil
}
