package crypt_security

import "github.com/urfave/cli"

var DigesterCMD = cli.Command{
	Name:     "digester",
	Aliases:  []string{"digest"},
	Usage:    "Computes a digest from a string using different algorithms.",
	Category: "Cryptography & Security",
	Action: func(c *cli.Context) error {
		return nil
	},
}
