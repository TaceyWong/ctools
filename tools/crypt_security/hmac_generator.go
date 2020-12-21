package crypt_security

import "github.com/urfave/cli/v2"

var HMACGeneratorCMD = cli.Command{
	Name:     "hmac_gen",
	Aliases:  []string{"hmac_g"},
	Usage:    "基于哈希消息验证的(HMAC)计算",
	Category: "加密安全",
	Action: func(c *cli.Context) error {
		return nil
	},
}
