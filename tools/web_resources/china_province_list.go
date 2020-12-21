package web_resources

import "github.com/urfave/cli"

var ChinaProvinceCMD = cli.Command{
	Name:     "china_province",
	Aliases:  []string{"chinap"},
	Usage:    " snippets to generate a list of provinces and territories for China",
	Category: "网页资源",
	Action: func(c *cli.Context) error {
		return nil
	},
}
