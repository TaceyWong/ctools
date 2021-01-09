package crypt_security

import "github.com/urfave/cli/v2"

var SHA512GeneratorCMD = cli.Command{
	Name:     "sha512_gen",
	Aliases:  []string{"sha512_g"},
	Usage:    "Computes a digest from a string using SHA-512. ",
	Category: "Cryptography & Security",
	Action: func(c *cli.Context) error {
		return nil
	},
}
