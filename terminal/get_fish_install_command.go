package terminal

func GetFishInstallCommand(distro string) (string, []string) {
	switch distro {
	case "OpenSUSE":
		return "zyper", []string{"addrepo", "https://download.opensuse.org/repositories/shells:fish:release:3/openSUSE_Tumbleweed/shells:fish:release:3.repo", "&&", "zypper refresh", "&&", "zypper install", "fish"}
	case "Arch Linux":
		return "pacman", []string{"-Syy", "fish"}
	case "Fedora":
		return "dnf", []string{"install", "fish"}
	}

	return "apt", []string{"install", "-y", "zsh"}

}
