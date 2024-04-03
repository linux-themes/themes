package commands

func SetCommand() {
	InDevelopment()

	// option := ""
	// if option == "icons" {
	// 	fmt.Println("icons")
	// }
	// if option == "themes" {
	// 	fmt.Println("themes")
	// }

	// envirnoment := ""
	// switch envirnoment {
	// case "gnome":
	// 	fmt.Println("gnome")
	// case "kde":
	// 	fmt.Println("kde")
	// case "xfce":
	// 	fmt.Println("xfce")
	// default:
	// 	fmt.Println("Set Error.")
	// 	os.Exit(0)
	// }

	// cmd := exec.Command("tar", "-xf")
	// _, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
}

// Gnome
// Change GTK-Theme:
// gsettings set org.gnome.desktop.interface gtk-theme "CoolestThemeOnEarth"
// Change Icon-Theme:
// gsettings set org.gnome.desktop.interface icon-theme 'MyIconTheme'
// Change Window-Theme:
// gsettings set org.gnome.desktop.wm.preferences theme "CoolestThemeOnEarth"

// KDE
// /usr/lib/plasma-changeicons breeze
// /usr/lib/plasma-changeicons breeze-dark
// /usr/lib/plasma-changeicons breath # Manjaro

// xfce
// xfconf-query -c xsettings -p /Net/IconThemeName -s elementaryXubuntu
// xfconf-query -c xsettings -p /Net/ThemeName -s "Clearlooks", same for /Net/IconThemeName.
