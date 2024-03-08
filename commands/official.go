package commands

type Package struct {
	Name string
	Path string
}

type Packages struct {
	Packages []Package
}

func ListOffical()     {}
func InstallSelected() {}
