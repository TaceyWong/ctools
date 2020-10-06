package web_resources

import "github.com/urfave/cli"

var TimezoneListCMD = cli.Command{
	Name:     "timezone_list",
	Aliases:  []string{"timezone_l"},
	Usage:    "snippets to generate a list of time zones",
	Category: "Web Resources",
	Action: func(c *cli.Context) error {
		return nil
	},
}
