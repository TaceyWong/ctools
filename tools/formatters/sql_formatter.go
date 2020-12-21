package formatters

import "github.com/urfave/cli/v2"

var SQLFormatCMD = cli.Command{
	Name:     "sql_format",
	Aliases:  []string{"sf"},
	Usage:    "格式化SQL",
	Category: "格式化器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
