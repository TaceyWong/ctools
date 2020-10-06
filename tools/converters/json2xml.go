package converters

import "github.com/urfave/cli"

var JSON2XMLCMD = cli.Command{
	Name:     "json2xml",
	Aliases:  []string{"c2x"},
	Usage:    "convert a JSON file into an xml file",
	Category: "Converters",
	Action: func(c *cli.Context) error {
		return nil
	},
}
