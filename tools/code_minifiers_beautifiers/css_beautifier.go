package code_minifiers_beautifiers

import "github.com/urfave/cli/v2"

var CSSBeautifierCMD = cli.Command{
	Name:     "css_beauty",
	Aliases:  []string{"css_p"},
	Usage:    "Formats your CSS files for optimal readability. ",
	Category: "Code Minifiers / Beautifier",
	Action: func(c *cli.Context) error {
		return nil
	},
}
