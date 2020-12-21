package code_minifiers_beautifiers

import "github.com/urfave/cli"

// JavascriptBeautifierCMD  js代码压缩命令
var JavascriptBeautifierCMD = cli.Command{
	Name:     "javascript_beauty",
	Aliases:  []string{"js_beauty"},
	Usage:    "格式化JS代码",
	Category: "代码压缩和美化",
	Action: func(c *cli.Context) error {
		return nil
	},
}
