package converters

import "github.com/urfave/cli/v2"

var CSV2JSONCMD = cli.Command{
	Name:     "csv2json",
	Aliases:  []string{"c2j"},
	Usage:    "CSV转换为JSON",
	Category: "转换工具",
	Action: func(c *cli.Context) error {
		return nil
	},
}
