package string_utils

import "github.com/urfave/cli"

var HTMLEscapeCMD = cli.Command{
	Name:     "html_escape",
	Aliases:  []string{"he"},
	Usage:    "Escapes or unescapes an HTML file removing traces of offending characters that could be wrongfully interpreted as markup.",
	Category: "字符串工具",
	Action: func(c *cli.Context) error {
		return nil
	},
}
