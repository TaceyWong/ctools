package web_resources

import "github.com/urfave/cli/v2"

var USAStateCMD = cli.Command{
	Name:     "usa_state",
	Aliases:  []string{"mexcio_s"},
	Usage:    "美国州列表代码",
	Category: "网页资源",
	Action: func(c *cli.Context) error {
		return nil
	},
}
