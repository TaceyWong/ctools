package string_utils

import "github.com/urfave/cli"

var XMLEscapeCMD = cli.Command{
	Name:     "xml_escape",
	Aliases:  []string{"xe"},
	Usage:    "Escapes or unescapes an XML file removing traces of offending characters that could be wrongfully interpreted as markup.",
	Category: "字符串工具",
	Action: func(c *cli.Context) error {
		return nil
	},
}
