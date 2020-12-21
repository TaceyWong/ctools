package validators

import "github.com/urfave/cli/v2"

var RegexTestCMD = cli.Command{
	Name:     "regex_test",
	Aliases:  []string{"rt"},
	Usage:    "高亮正则匹配",
	Category: "验证器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
