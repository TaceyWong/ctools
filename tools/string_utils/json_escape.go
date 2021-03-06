package string_utils

import "github.com/urfave/cli/v2"

var JSONEscapeCMD = cli.Command{
	Name:     "json_escape",
	Aliases:  []string{"jsone"},
	Usage:    "Escapes or unescapes a JSON string removing traces of offending characters that could prevent parsing.",
	Category: "字符串工具",
	Action: func(c *cli.Context) error {
		return nil
	},
}
