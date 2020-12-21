package crypt_security

import "github.com/urfave/cli/v2"

var SHA256GeneratorCMD = cli.Command{
	Name:     "sha256_gen",
	Aliases:  []string{"sha256_g"},
	Usage:    "生成字符串SHA-256. ",
	Category: "加密安全",
	Action: func(c *cli.Context) error {
		return nil
	},
}
