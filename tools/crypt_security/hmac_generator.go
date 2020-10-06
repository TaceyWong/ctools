package crypt_security

import "github.com/urfave/cli"

var HMACGeneratorCMD = cli.Command{
	Name:     "hmac_gen",
	Aliases:  []string{"hmac_g"},
	Usage:    "Computes a Hash-based message authentication code (HMAC) using different algorithms.",
	Category: "Cryptography & Security",
	Action: func(c *cli.Context) error {
		return nil
	},
}
