package utils

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func SendDistroSurvey() (string, error) {
	qs := []*survey.Question{
		{
			Name: "distro",
			Prompt: &survey.Select{
				Message: "Select your distro:",
				Options: []string{"Ubuntu", "OpenSUSE", "Arch Linux", "Void Linux", "Fedora", "OpenBSD"},
				Default: "Ubuntu",
			},
		},
	}

	answers := struct {
		Distro string `survey:"distro"`
	}{}

	fmt.Println("Detected Linux OS. Please select your distro:")
	err := survey.Ask(qs, &answers)

	if err != nil {
		return "", err
	}

	return answers.Distro, nil
}
