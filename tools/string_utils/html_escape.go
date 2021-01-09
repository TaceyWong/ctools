package string_utils

import "github.com/urfave/cli/v2"

var HTMLEscapeCMD = cli.Command{
	Name:     "html_escape",
	Aliases:  []string{"he"},
	Usage:    "Escapes or unescapes an HTML file removing traces of offending characters that could be wrongfully interpreted as markup.",
	Category: "String Escaper & Utilities",
	Action: func(c *cli.Context) error {
		return nil
	},
}
