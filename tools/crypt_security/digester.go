package crypt_security

import "github.com/urfave/cli/v2"

var DigesterCMD = cli.Command{
	Name:     "digester",
	Aliases:  []string{"digest"},
	Usage:    "使用不同算法从字符串计算一个数字",
	Category: "加密安全",
	Action: func(c *cli.Context) error {
		return nil
	},
}
