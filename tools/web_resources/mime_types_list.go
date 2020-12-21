package web_resources

import "github.com/urfave/cli"

var MIMETypeListCMD = cli.Command{
	Name:     "mime_type_list",
	Aliases:  []string{"mtl"},
	Usage:    "Full list of MIME types",
	Category: "网页资源",
	Action: func(c *cli.Context) error {
		return nil
	},
}
