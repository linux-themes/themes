package commands

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
)

var (
	red    = lipgloss.AdaptiveColor{Light: "#FE5F86", Dark: "#FE5F86"}
	indigo = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green  = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
)

type Styles struct {
	Base,
	HeaderText,
	Status,
	StatusHeader,
	Highlight,
	ErrorHeaderText,
	Help lipgloss.Style
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().
		Padding(1, 4, 0, 1)
	s.HeaderText = lg.NewStyle().
		Foreground(indigo).
		Bold(true).
		Padding(0, 1, 0, 2)
	s.Status = lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(indigo).
		PaddingLeft(1).
		MarginTop(1)
	s.StatusHeader = lg.NewStyle().
		Foreground(green).
		Bold(true)
	s.Highlight = lg.NewStyle().
		Foreground(lipgloss.Color("212"))
	s.ErrorHeaderText = s.HeaderText.Copy().
		Foreground(red)
	s.Help = lg.NewStyle().
		Foreground(lipgloss.Color("240"))
	return &s
}

// type state int

// const (
// 	statusNormal state = iota
// 	stateDone
// )

// type Model struct {
// 	state  state
// 	lg     *lipgloss.Renderer
// 	styles *Styles
// 	form   *huh.Form
// 	width  int
// }

type Option int

const (
	Icons Option = iota + 1
	Themes
	Configs
)

func (option Option) String() string {
	switch option {
	case Icons:
		return ".icons"
	case Themes:
		return ".themes"
	case Configs:
		return ".config"
	default:
		return ""
	}
}

type Command struct {
	Type     string
	Option   Option
	Packages []string
}

func RemoveCommand() {
	var command Command
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions(ICONS, THEMES, CONFIG)...).
				Title("Linux Themes Remover").
				Value(&command.Type),
		),
	).WithAccessible(accessible).WithTheme(huh.ThemeDracula())
	err := form.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	packages := GetPackages(command.Type)
	total_packages := len(packages)
	str_builder := "Installed Themes: " + strconv.Itoa(total_packages)
	options := []huh.Option[string]{}
	for index := 0; index < len(packages); index++ {
		option := huh.Option[string]{}
		option.Key = packages[index]
		option.Value = packages[index]
		option.Selected(false)
		options = append(options, option)
	}
	form2 := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Options(options...).
				Title(str_builder).
				Value(&command.Packages),
		),
	).WithAccessible(accessible).WithTheme(huh.ThemeDracula())
	err = form2.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	loading_action_progress := func() {
		// time.Sleep(1 * time.Second)
		Remove(command.Packages)
		fmt.Println(command.Packages)
	}
	_ = spinner.New().
		Title("Removing Themes...").
		TitleStyle(lipgloss.NewStyle().Foreground(green)).
		Style(lipgloss.NewStyle().Foreground(red)).
		Accessible(accessible).
		Action(loading_action_progress).
		Run()

	fmt.Println(lipgloss.NewStyle().Foreground(green).SetString("Themes removed: " + strconv.Itoa(len(command.Packages))))
	for _, packag := range command.Packages {
		fmt.Println(packag)
	}

}

func GetPackages(category string) []string {
	var path string
	if category == ".icons" {
		path = ICON_PATH
	}
	if category == ".themes" {
		path = THEME_PATH
	}
	if category == ".config" {
		path = CONFIG_PATH
	}

	packages := []string{}
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			packages = append(packages, path+"/"+file.Name())
		}
	}
	return packages
}

func Remove(packages []string) {
	for _, packag := range packages {
		fmt.Print("Removing: ")
		fmt.Println(packag)
		if err := os.RemoveAll(packag); err != nil {
			fmt.Println(err.Error())
		}
	}
}
