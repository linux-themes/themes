package cmd

import "github.com/charmbracelet/lipgloss"

const (
	BLUE      = "\033[1;34m"
	BLUE_THIN = "\033[0;36m"
	CYAN      = "\033[1;36m"
	YELLOW    = "\033[1;33m"
	RED       = "\033[1;31ms"
	GREEN     = "\033[1;32m"
	RESET     = "\033[0m"
	FUCHSIA   = lipgloss.Color("#F780E2")
	YELLOW_L  = lipgloss.ANSIColor(11)
)

var (
	normalFg = lipgloss.AdaptiveColor{Light: "235", Dark: "252"}
	indigo   = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	cream    = lipgloss.AdaptiveColor{Light: "#FFFDF5", Dark: "#FFFDF5"}
	green    = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
	red      = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}
)
