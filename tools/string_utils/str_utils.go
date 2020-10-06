package string_utils

import "github.com/urfave/cli"

var StrUtilCMD = cli.Command{
	Name:     "str_util",
	Aliases:  []string{"su"},
	Usage:    "A couple of online string utilities written in JavaScript.",
	Category: "String Escaper & Utilities",
	Action: func(c *cli.Context) error {
		return nil
	},
}
