package formatters

import "github.com/urfave/cli"

var SQLFormatCMD = cli.Command{
	Name:     "sql_format",
	Aliases:  []string{"sf"},
	Usage:    "Format SQL Text",
	Category: "Formatters",
	Action: func(c *cli.Context) error {
		return nil
	},
}
