package converters

import "github.com/urfave/cli/v2"

var JSON2GoCMD = cli.Command{
	Name:     "json2go",
	Aliases:  []string{"c2g"},
	Usage:    "JSON转换为go结构体",
	Category: "转换器",
	Action: func(c *cli.Context) error {

		return nil
	},
}
