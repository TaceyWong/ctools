package converters

import "github.com/urfave/cli/v2"

var XML2JSONCMD = cli.Command{
	Name:     "xml2json",
	Aliases:  []string{"x2j"},
	Usage:    "XML转换为JSON",
	Category: "转换器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
