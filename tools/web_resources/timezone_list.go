package web_resources

import "github.com/urfave/cli/v2"

var TimezoneListCMD = cli.Command{
	Name:     "timezone_list",
	Aliases:  []string{"timezone_l"},
	Usage:    "时区列表代码",
	Category: "网页资源",
	Action: func(c *cli.Context) error {
		return nil
	},
}
