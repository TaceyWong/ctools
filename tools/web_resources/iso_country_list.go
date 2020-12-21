package web_resources

import "github.com/urfave/cli/v2"

var ISOCountryListCMD = cli.Command{
	Name:     "iso_country_list",
	Aliases:  []string{"icl"},
	Usage:    "国家列表代码(ISO-3166-1 & ISO-3166-2)",
	Category: "网页资源",
	Action: func(c *cli.Context) error {
		return nil
	},
}
