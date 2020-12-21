package web_resources

import "github.com/urfave/cli/v2"

var LessCompilerCMD = cli.Command{
	Name:     "less_compiler",
	Aliases:  []string{"less_c"},
	Usage:    "LESS编译为CSS",
	Category: "网页资源",
	Action: func(c *cli.Context) error {
		return nil
	},
}
