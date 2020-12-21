package formatters

import "github.com/urfave/cli"

var HTMLFormatCMD = cli.Command{
	Name:     "html_format",
	Aliases:  []string{"hf"},
	Usage:    "Format HTML Text",
	Category: "格式化器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
