package web_resources

import "github.com/urfave/cli/v2"

var ChinaProvinceCMD = cli.Command{
	Name:     "china_province",
	Aliases:  []string{"chinap"},
	Usage:    " snippets to generate a list of provinces and territories for China",
	Category: "Web Resources",
	Action: func(c *cli.Context) error {
		return nil
	},
}
