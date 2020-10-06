package web_resources

import "github.com/urfave/cli"

var USAStateCMD = cli.Command{
	Name:     "usa_state",
	Aliases:  []string{"mexcio_s"},
	Usage:    "Snippets to generate a list of states for the United States",
	Category: "Web Resources",
	Action: func(c *cli.Context) error {
		return nil
	},
}
