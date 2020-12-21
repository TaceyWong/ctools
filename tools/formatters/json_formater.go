package formatters

import "github.com/urfave/cli/v2"

var JSONFormatCMD = cli.Command{
	Name:     "json_format",
	Aliases:  []string{"jf"},
	Usage:    "格式化JSON",
	Category: "格式化器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
