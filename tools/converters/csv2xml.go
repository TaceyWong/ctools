package converters

import "github.com/urfave/cli/v2"

var CSV2XMLCMD = cli.Command{
	Name:     "csv2xml",
	Aliases:  []string{"c2x"},
	Usage:    "convert a CSV file into an XML file",
	Category: "Converters",
	Action: func(c *cli.Context) error {
		return nil
	},
}
