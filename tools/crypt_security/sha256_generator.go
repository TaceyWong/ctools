package crypt_security

import "github.com/urfave/cli/v2"

var SHA256GeneratorCMD = cli.Command{
	Name:     "sha256_gen",
	Aliases:  []string{"sha256_g"},
	Usage:    "Computes a digest from a string using SHA-256. ",
	Category: "Cryptography & Security",
	Action: func(c *cli.Context) error {
		return nil
	},
}
