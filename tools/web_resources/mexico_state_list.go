package web_resources

import "github.com/urfave/cli/v2"

var MexcioStateCMD = cli.Command{
	Name:     "mexico_state",
	Aliases:  []string{"mexcio_s"},
	Usage:    "snippets to generate a list state for Canada",
	Category: "Web Resources",
	Action: func(c *cli.Context) error {
		return nil
	},
}
