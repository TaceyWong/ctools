package converters

import "github.com/urfave/cli/v2"

var Timestamp2DateCMD = cli.Command{
	Name:     "timestamp2date",
	Aliases:  []string{"t2d"},
	Usage:    "epoch/unix时间戳转换为人类可读日期",
	Category: "转换器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
