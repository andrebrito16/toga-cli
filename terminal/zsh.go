package terminal

import (
	"fmt"

	"github.com/andrebrito16/toga-cli/evaluator"
	"github.com/andrebrito16/toga-cli/utils"
	"github.com/urfave/cli/v2"
)

func ZshInstall(c *cli.Context) error {
	os := utils.DetectOs()

	switch os {
	case "linux":
		return ZshInstallLinux()
	case "mac":
		return ZshInstallMac()
	}

	return nil
}

func ZshInstallLinux() error {
	distro, err := utils.SendDistroSurvey()

	if err != nil {
		return err
	}

	installCommand, installArgs := GetZshInstallCommand(distro)

	err = evaluator.Run(installCommand, installArgs)

	if err != nil {
		return err
	}

	return nil
}

func ZshInstallMac() error {
	err := evaluator.Run("brew", []string{"version"})

	if err != nil {
		fmt.Println("Brew not found. Please install it first.")
		return err
	}

	err = evaluator.Run("brew", []string{"install", "zsh"})

	if err != nil {
		return err
	}

	return nil
}
