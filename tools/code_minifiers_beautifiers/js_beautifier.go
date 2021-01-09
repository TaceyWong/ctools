package code_minifiers_beautifiers

import "github.com/urfave/cli/v2"

var JavascriptBeautifierCMD = cli.Command{
	Name:     "javascript_beauty",
	Aliases:  []string{"js_beauty"},
	Usage:    "Formats your JS files for optimal readability",
	Category: "Code Minifiers / Beautifier",
	Action: func(c *cli.Context) error {
		return nil
	},
}
