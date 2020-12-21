package crypt_security

import "github.com/urfave/cli/v2"

var SHA512GeneratorCMD = cli.Command{
	Name:     "sha512_gen",
	Aliases:  []string{"sha512_g"},
	Usage:    "生成字符串SHA-512. ",
	Category: "加密安全",
	Action: func(c *cli.Context) error {
		return nil
	},
}
