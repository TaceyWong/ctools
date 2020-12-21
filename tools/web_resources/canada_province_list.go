package web_resources

import "github.com/urfave/cli/v2"

var CanadaProvinceCMD = cli.Command{
	Name:     "canada_province",
	Aliases:  []string{"cprovince"},
	Usage:    "加拿大省份列表代码块儿",
	Category: "网页资源",
	Action: func(c *cli.Context) error {
		return nil
	},
}
