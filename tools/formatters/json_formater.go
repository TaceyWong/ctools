package formatters

import "github.com/urfave/cli/v2"

var JSONFormatCMD = cli.Command{
	Name:     "json_format",
	Aliases:  []string{"jf"},
	Usage:    "Format JSON Text",
	Category: "Formatters",
	Action: func(c *cli.Context) error {
		return nil
	},
}
