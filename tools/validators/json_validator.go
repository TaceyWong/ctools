package validators

import "github.com/urfave/cli/v2"

var JSONValidCMD = cli.Command{
	Name:     "json_valid",
	Aliases:  []string{"jv"},
	Usage:    "验证JSON字符串是否符合RFC 4627",
	Category: "验证器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
