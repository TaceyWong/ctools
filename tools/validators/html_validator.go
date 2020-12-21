package validators

import "github.com/urfave/cli/v2"

var HTMLValidCMD = cli.Command{
	Name:     "html_valid",
	Aliases:  []string{"hv"},
	Usage:    "验证HTML文本是否符合&兼容w3c标注",
	Category: "验证器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
