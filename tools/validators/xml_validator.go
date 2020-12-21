package validators

import "github.com/urfave/cli/v2"

var XMLValidCMD = cli.Command{
	Name:     "xml_valid",
	Aliases:  []string{"xv"},
	Usage:    "验证XML是否符合XSD规范",
	Category: "验证器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
