package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func SetCommand(category string) {
	InDevelopment()
	// switch category {
	// case ICONS:
	// 	Set(category)
	// case THEMES:
	// 	Set(category)
	// default:
	// 	HelpCommand()
	// }
}

func Set(category string) {
	desktop_enviroment := os.Getenv("XDG_SESSION_DESKTOP")
	fmt.Println(desktop_enviroment)

	icon_variable := "gnome"
	theme_variable := "gnome"

	switch desktop_enviroment {
	case "gnome":
		if category == ICONS {
			cmd := exec.Command("gsettings",
				"set",
				"org.gnome.desktop.interface",
				"icon-theme",
				icon_variable)
			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			}
		}
		if category == THEMES {
			cmd := exec.Command("gsettings",
				"set",
				"org.gnome.shell.extensions.user-theme",
				"name",
				theme_variable)
			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			}
		}
	case "kde":
		fmt.Println("kde")
	default:
		fmt.Println("Set Error")
	}
}

// user theme extensions must be enabled - add check
// if 'variable' does not exist default 'Gnome' is set.

// Gnome
// Change Icon-Theme:
// gsettings set org.gnome.desktop.interface icon-theme 'Gnome'
// gsettings set org.gnome.desktop.interface icon-theme 'gicons'
// Change Shell-Theme:
// gsettings set org.gnome.shell.extensions.user-theme name "Zukitwo"

// KDE
// /usr/lib/plasma-changeicons breeze
// /usr/lib/plasma-changeicons breeze-dark
// /usr/lib/plasma-changeicons breath # Manjaro

// xfce
// xfconf-query -c xsettings -p /Net/IconThemeName -s elementaryXubuntu
// xfconf-query -c xsettings -p /Net/ThemeName -s "Clearlooks", same for /Net/IconThemeName.
