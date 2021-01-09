package validators

import "github.com/urfave/cli/v2"

var CronGenCMD = cli.Command{
	Name:     "cron_gen",
	Aliases:  []string{"cg"},
	Usage:    "Generate a quartz cron expression with an easy to use interface",
	Category: "Validators",
	Action: func(c *cli.Context) error {
		return nil
	},
}
