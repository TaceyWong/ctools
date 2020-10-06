package converters

import "github.com/urfave/cli"

var CSV2JSONCMD = cli.Command{
	Name:     "csv2json",
	Aliases:  []string{"c2j"},
	Usage:    "convert a CSV file into an JSON file",
	Category: "Converters",
	Action: func(c *cli.Context) error {
		return nil
	},
}
