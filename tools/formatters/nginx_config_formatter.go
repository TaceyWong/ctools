package formatters

import "github.com/urfave/cli/v2"

var NginxConfigFormatCMD = cli.Command{
	Name:     "nginx_conf_format",
	Aliases:  []string{"nginx_conf_f"},
	Usage:    "格式化Nginx配置",
	Category: "格式化器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
