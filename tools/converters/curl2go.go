package converters

import "github.com/urfave/cli/v2"

var CURL2GoCMD = cli.Command{
	Name:     "curl2go",
	Aliases:  []string{"curl2g"},
	Usage:    "curl转换为Go代码",
	Category: "转换工具",
	Action: func(c *cli.Context) error {
		return nil
	},
}
