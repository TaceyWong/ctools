package converters

import "github.com/urfave/cli/v2"

var CSV2XMLCMD = cli.Command{
	Name:     "csv2xml",
	Aliases:  []string{"c2x"},
	Usage:    "CSV转换为xml",
	Category: "转换工具",
	Action: func(c *cli.Context) error {
		return nil
	},
}
