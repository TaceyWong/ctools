package crypt_security

import "github.com/urfave/cli/v2"

var MD5GeneratorCMD = cli.Command{
	Name:     "md5_gen",
	Aliases:  []string{"md5_g"},
	Usage:    "生成字符串MD5",
	Category: "加密安全",
	Action: func(c *cli.Context) error {
		return nil
	},
}
