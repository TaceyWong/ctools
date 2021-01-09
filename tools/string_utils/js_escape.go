package string_utils

import "github.com/urfave/cli/v2"

var JavascriptEscapeCMD = cli.Command{
	Name:     "js_escape",
	Aliases:  []string{"je"},
	Usage:    "Escapes or unescapes a JavaScript string removing traces of offending characters that could prevent interpretation.",
	Category: "String Escaper & Utilities",
	Action: func(c *cli.Context) error {
		return nil
	},
}
