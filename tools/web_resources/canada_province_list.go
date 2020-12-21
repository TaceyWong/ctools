package web_resources

import "github.com/urfave/cli"

var CanadaProvinceCMD = cli.Command{
	Name:     "canada_province",
	Aliases:  []string{"cprovince"},
	Usage:    " snippets to generate a list of provinces and territories for Canada",
	Category: "网页资源",
	Action: func(c *cli.Context) error {
		return nil
	},
}
