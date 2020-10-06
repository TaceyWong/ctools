package string_utils

import "github.com/urfave/cli"

var JSONEscapeCMD = cli.Command{
	Name:     "json_escape",
	Aliases:  []string{"jsone"},
	Usage:    "Escapes or unescapes a JSON string removing traces of offending characters that could prevent parsing.",
	Category: "String Escaper & Utilities",
	Action: func(c *cli.Context) error {
		return nil
	},
}
