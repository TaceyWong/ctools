package formatters

import "github.com/urfave/cli"

var XMLFormatCMD = cli.Command{
	Name:     "xml_format",
	Aliases:  []string{"hf"},
	Usage:    "Format XML Text",
	Category: "Formatters",
	Action: func(c *cli.Context) error {
		return nil
	},
}
