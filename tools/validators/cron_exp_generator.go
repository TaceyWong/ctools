package validators

import "github.com/urfave/cli/v2"

var CronGenCMD = cli.Command{
	Name:     "cron_gen",
	Aliases:  []string{"cg"},
	Usage:    "生成quartz cron 表达式的简单用户接口",
	Category: "验证器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
