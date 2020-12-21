package converters

import "github.com/urfave/cli"

var CURL2GoCMD = cli.Command{
	Name:     "curl2go",
	Aliases:  []string{"curl2g"},
	Usage:    " turns a curl command into Go code.",
	Category: "转换器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
