package code_minifiers_beautifiers

import "github.com/urfave/cli/v2"

var CSSMinifierCMD = cli.Command{
	Name:     "css_mini",
	Aliases:  []string{"css_m"},
	Usage:    "Compresses a CSS string/file with no possible side-effect",
	Category: "Code Minifiers / Beautifier",
	Action: func(c *cli.Context) error {
		return nil
	},
}
