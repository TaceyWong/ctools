package converters

import "github.com/urfave/cli"

var JSON2GoCMD = cli.Command{
	Name:     "json2go",
	Aliases:  []string{"c2g"},
	Usage:    "convert a JSON file into an go struct file",
	Category: "转换器",
	Action: func(c *cli.Context) error {

		return nil
	},
}
