package terminal

import (
	"fmt"

	"github.com/andrebrito16/toga-cli/evaluator"
	"github.com/andrebrito16/toga-cli/utils"
	"github.com/urfave/cli/v2"
)

func FishInstall(c *cli.Context) error {
	os := utils.DetectOs()

	switch os {
	case "linux":
		return FishInstallLinux()
	case "mac":
		return FishInstallMac()
	}

	return nil
}

func FishInstallLinux() error {
	distro, err := utils.SendDistroSurvey()

	if err != nil {
		return err
	}

	if distro == "Void Linux" {
		fmt.Printf("Fish is not available for %s. Please install it manually.", distro)
		return nil
	}

	installCommand, installArgs := GetFishInstallCommand(distro)

	err = evaluator.Run(installCommand, installArgs)

	if err != nil {
		return err
	}

	return nil
}

func FishInstallMac() error {
	err := evaluator.Run("brew", []string{"version"})

	if err != nil {
		fmt.Println("Brew not found. Please install it first.")
		return err
	}

	err = evaluator.Run("brew", []string{"install", "fish"})

	if err != nil {
		return err
	}

	return nil
}
