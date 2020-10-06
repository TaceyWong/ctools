package string_utils

import "github.com/urfave/cli"

var XMLEscapeCMD = cli.Command{
	Name:     "xml_escape",
	Aliases:  []string{"xe"},
	Usage:    "Escapes or unescapes an XML file removing traces of offending characters that could be wrongfully interpreted as markup.",
	Category: "String Escaper & Utilities",
	Action: func(c *cli.Context) error {
		return nil
	},
}
