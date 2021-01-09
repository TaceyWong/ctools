package converters

import "github.com/urfave/cli/v2"

var XML2JSONCMD = cli.Command{
	Name:     "xml2json",
	Aliases:  []string{"x2j"},
	Usage:    "convert a XML file into an JSON file",
	Category: "Converters",
	Action: func(c *cli.Context) error {
		return nil
	},
}
