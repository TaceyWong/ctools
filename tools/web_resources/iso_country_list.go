package web_resources

import "github.com/urfave/cli"

var ISOCountryListCMD = cli.Command{
	Name:     "iso_country_list",
	Aliases:  []string{"icl"},
	Usage:    " snippets to generate a list of countries using the ISO-3166-1 and ISO-3166-2 codes.",
	Category: "网页资源",
	Action: func(c *cli.Context) error {
		return nil
	},
}
