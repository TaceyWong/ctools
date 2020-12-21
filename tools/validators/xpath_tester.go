package validators

import "github.com/urfave/cli/v2"

var XPathTestCMD = cli.Command{
	Name:     "xpath_test",
	Aliases:  []string{"xt"},
	Usage:    "对XML文件执行XPath查询",
	Category: "验证器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
