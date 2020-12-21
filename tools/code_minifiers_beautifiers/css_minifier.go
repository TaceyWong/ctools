package code_minifiers_beautifiers

import "github.com/urfave/cli"

// CSSMinifierCMD css压缩命令
var CSSMinifierCMD = cli.Command{
	Name:     "css_mini",
	Aliases:  []string{"css_m"},
	Usage:    "无副作用压缩哟CSS代码",
	Category: "代码压缩和美化",
	Action: func(c *cli.Context) error {
		return nil
	},
}
