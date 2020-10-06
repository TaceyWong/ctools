package formatters

import "github.com/urfave/cli"

var NginxConfigFormatCMD = cli.Command{
	Name:     "nginx_conf_format",
	Aliases:  []string{"nginx_conf_f"},
	Usage:    "Format HTML Text",
	Category: "Formatters",
	Action: func(c *cli.Context) error {
		return nil
	},
}
