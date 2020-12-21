package crypt_security

import "github.com/urfave/cli"

var MD5GeneratorCMD = cli.Command{
	Name:     "md5_gen",
	Aliases:  []string{"md5_g"},
	Usage:    "Computes a digest from a string using MD5",
	Category: "加密安全",
	Action: func(c *cli.Context) error {
		return nil
	},
}
