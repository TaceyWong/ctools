package code_minifiers_beautifiers

import "github.com/urfave/cli/v2"

var JavascriptMinifierCMD = cli.Command{
	Name:     "javascript_mini",
	Aliases:  []string{"js_mini"},
	Usage:    "Compresses a JavaScript string/file with no possible side-effect",
	Category: "Code Minifiers / Beautifier",
	Action: func(c *cli.Context) error {
		return nil
	},
}
