package converters

import "github.com/urfave/cli/v2"

var JSON2GoCMD = cli.Command{
	Name:     "json2go",
	Aliases:  []string{"c2g"},
	Usage:    "convert a JSON file into an go struct file",
	Category: "Converters",
	Action: func(c *cli.Context) error {

		return nil
	},
}
