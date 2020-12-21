package formatters

import "github.com/urfave/cli/v2"

var XMLFormatCMD = cli.Command{
	Name:     "xml_format",
	Aliases:  []string{"hf"},
	Usage:    "格式化XML",
	Category: "格式化器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
