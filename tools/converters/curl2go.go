package converters

import "github.com/urfave/cli"

var CURL2GoCMD = cli.Command{
	Name:     "curl2go",
	Aliases:  []string{"curl2g"},
	Usage:    " turns a curl command into Go code.",
	Category: "Converters",
	Action: func(c *cli.Context) error {
		return nil
	},
}