package string_utils

import "github.com/urfave/cli/v2"

var CSVEscapeCMD = cli.Command{
	Name:     "csv_escape",
	Aliases:  []string{"ce"},
	Usage:    "Escapes or unescapes a CSV string removing traces of offending characters that could prevent parsing.",
	Category: "字符串工具",
	Action: func(c *cli.Context) error {
		return nil
	},
}
