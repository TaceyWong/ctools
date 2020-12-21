package formatters

import "github.com/urfave/cli/v2"

var HTMLFormatCMD = cli.Command{
	Name:     "html_format",
	Aliases:  []string{"hf"},
	Usage:    "格式化HTML",
	Category: "格式化器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
