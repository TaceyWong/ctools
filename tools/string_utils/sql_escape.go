package string_utils

import "github.com/urfave/cli/v2"

var SQLEscapeCMD = cli.Command{
	Name:    "sql_escape",
	Aliases: []string{"sqle"},
	Usage: "Escapes or unescapes a SQL string removing traces of offending characters that could prevent execution.	",
	Category: "String Escaper & Utilities",
	Action: func(c *cli.Context) error {
		return nil
	},
}
