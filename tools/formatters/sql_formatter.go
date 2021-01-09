package formatters

import "github.com/urfave/cli/v2"

var SQLFormatCMD = cli.Command{
	Name:     "sql_format",
	Aliases:  []string{"sf"},
	Usage:    "Format SQL Text",
	Category: "Formatters",
	Action: func(c *cli.Context) error {
		return nil
	},
}
