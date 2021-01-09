package converters

import "github.com/urfave/cli/v2"

var Timestamp2DateCMD = cli.Command{
	Name:     "timestamp2date",
	Aliases:  []string{"t2d"},
	Usage:    "Converts an epoch/unix timestamp into a human readable date",
	Category: "Converters",
	Action: func(c *cli.Context) error {
		return nil
	},
}
