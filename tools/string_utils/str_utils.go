package string_utils

import "github.com/urfave/cli"

var StrUtilCMD = cli.Command{
	Name:     "str_util",
	Aliases:  []string{"su"},
	Usage:    "A couple of online string utilities written in JavaScript.",
	Category: "字符串工具",
	Action: func(c *cli.Context) error {
		return nil
	},
}
