package converters

import "github.com/urfave/cli/v2"

var CURL2GoCMD = cli.Command{
	Name:     "curl2go",
	Aliases:  []string{"curl2g"},
	Usage:    "curl命令转换为Go代码",
	Category: "转换器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
