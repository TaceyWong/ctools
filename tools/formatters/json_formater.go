package formatters

import "github.com/urfave/cli"

var JSONFormatCMD = cli.Command{
	Name:     "json_format",
	Aliases:  []string{"jf"},
	Usage:    "Format JSON Text",
	Category: "格式化器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
