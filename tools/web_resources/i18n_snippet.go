package web_resources

import "github.com/urfave/cli/v2"

var I18nCMD = cli.Command{
	Name:     "i18n",
	Aliases:  []string{"i18"},
	Usage:    "日期时间货币等列表代码",
	Category: "网页资源",
	Action: func(c *cli.Context) error {
		return nil
	},
}
