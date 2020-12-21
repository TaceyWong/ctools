package converters

import "github.com/urfave/cli"

var CSV2XMLCMD = cli.Command{
	Name:     "csv2xml",
	Aliases:  []string{"c2x"},
	Usage:    "convert a CSV file into an XML file",
	Category: "转换器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
