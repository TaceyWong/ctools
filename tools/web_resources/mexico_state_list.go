package web_resources

import "github.com/urfave/cli/v2"

var MexcioStateCMD = cli.Command{
	Name:     "mexico_state",
	Aliases:  []string{"mexcio_s"},
	Usage:    "墨西哥州列表代码",
	Category: "网页资源",
	Action: func(c *cli.Context) error {
		return nil
	},
}
