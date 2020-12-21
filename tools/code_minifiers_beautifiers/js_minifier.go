package code_minifiers_beautifiers

import "github.com/urfave/cli"

// JavascriptMinifierCMD JS代码压缩命令
var JavascriptMinifierCMD = cli.Command{
	Name:     "javascript_mini",
	Aliases:  []string{"js_mini"},
	Usage:    "无副作用压缩JS代码",
	Category: "代码压缩和美化",
	Action: func(c *cli.Context) error {
		return nil
	},
}
