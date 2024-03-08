package markdown

import _ "embed"

//go:embed config.md
var CONFIG_MARKDOWN string

//go:embed help.md
var HELP_MARKDOWN string

//go:embed icons.md
var ICONS_MARKDOWN string

//go:embed contribute.md
var CONTRIBUTE_MARKDOWN string

//go:embed themes.md
var THEMES_MARKDOWN string
