package code_minifiers_beautifiers

import "github.com/urfave/cli/v2"

// CSSBeautifierCMD css_beauty command
var CSSBeautifierCMD = cli.Command{
	Name:     "css_beauty",
	Aliases:  []string{"css_p"},
	Usage:    "格式化CSS代码",
	Category: "代码压缩和美化",
	Action: func(c *cli.Context) error {
		return nil
	},
}
