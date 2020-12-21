package string_utils

import "github.com/urfave/cli/v2"

var JavascriptEscapeCMD = cli.Command{
	Name:     "js_escape",
	Aliases:  []string{"je"},
	Usage:    "Escapes or unescapes a JavaScript string removing traces of offending characters that could prevent interpretation.",
	Category: "字符串工具",
	Action: func(c *cli.Context) error {
		return nil
	},
}
