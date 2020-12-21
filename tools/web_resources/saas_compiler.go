package web_resources

import "github.com/urfave/cli/v2"

var SassCompilerCMD = cli.Command{
	Name:     "sass_compiler",
	Aliases:  []string{"sass_c"},
	Usage:    "SASS编译为CSS",
	Category: "网页资源",
	Action: func(c *cli.Context) error {
		return nil
	},
}
