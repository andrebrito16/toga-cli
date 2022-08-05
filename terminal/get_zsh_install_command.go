package terminal

func GetZshInstallCommand(distro string) (string, []string) {
	switch distro {
	case "OpenSUSE":
		return "zyper", []string{"install", "zsh"}
	case "Arch Linux":
		return "pacman", []string{"-Syy", "zsh"}
	case "Void Linux":
		return "xbps-install", []string{"zsh"}
	case "Fedora":
		return "dnf", []string{"install", "zsh"}
	}

	return "apt", []string{"install", "-y", "zsh"}
}
