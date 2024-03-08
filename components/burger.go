package components

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
)

type Option int

const (
	ICONS Option = iota + 1
	THEMES
	CONFIG
)

func (option Option) String() string {
	switch option {
	case ICONS:
		return ".icons"
	case THEMES:
		return ".themes"
	case CONFIG:
		return ".config"
	default:
		return ""
	}
}

type Order struct {
	Burger Burger
	Name   string
}

type Burger struct {
	Type   string
	Option Option
}

func MenuSelection() {
	var burger Burger
	var order = Order{Burger: burger}
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions(".icons", ".themes", ".config")...).
				Title("Linux Themes Remover").
				Value(&order.Burger.Type),
		),

		// get folders slice
		// map to options
		// make dynamic opiton
		huh.NewGroup(
			huh.NewSelect[Option]().
				Title("Select option").
				Options(
					huh.NewOption(".icons", ICONS).Selected(true),
					huh.NewOption(".themes", THEMES),
					huh.NewOption(".config", CONFIG),
				).
				Value(&order.Burger.Option),
		),
	).WithAccessible(accessible)
	err := form.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	prepareBurger := func() {
		time.Sleep(1 * time.Second)
	}
	_ = spinner.New().Title("Preparing your burger...").Accessible(accessible).Action(prepareBurger).Run()

	// Print results
	{
		var sb strings.Builder
		fmt.Fprintf(&sb,
			"%s\n\n",
			lipgloss.NewStyle().Bold(true).Render("BURGER RECEIPT"),
		)
		fmt.Fprintf(&sb, "\n\nThanks for your order!")
		fmt.Println(
			lipgloss.NewStyle().
				Width(40).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render(sb.String()),
		)
	}
}
