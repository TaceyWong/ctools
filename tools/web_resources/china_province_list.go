package web_resources

import "github.com/urfave/cli/v2"

var ChinaProvinceCMD = cli.Command{
	Name:     "china_province",
	Aliases:  []string{"chinap"},
	Usage:    "中国省份列表代码块儿",
	Category: "网页资源",
	Action: func(c *cli.Context) error {
		return nil
	},
}
