package web_resources

import "github.com/urfave/cli/v2"

var I18nCMD = cli.Command{
	Name:     "i18n",
	Aliases:  []string{"i18"},
	Usage:    " list of formatting standards for dates, time, decimals, currencies and more with code snippets ",
	Category: "Web Resources",
	Action: func(c *cli.Context) error {
		return nil
	},
}
