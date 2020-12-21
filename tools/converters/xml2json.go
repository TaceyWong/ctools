package converters

import "github.com/urfave/cli"

var XML2JSONCMD = cli.Command{
	Name:     "xml2json",
	Aliases:  []string{"x2j"},
	Usage:    "convert a XML file into an JSON file",
	Category: "转换器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
