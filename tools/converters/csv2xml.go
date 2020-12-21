package converters

import "github.com/urfave/cli/v2"

var CSV2XMLCMD = cli.Command{
	Name:     "csv2xml",
	Aliases:  []string{"c2x"},
	Usage:    "CSV转换为XML",
	Category: "转换器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
