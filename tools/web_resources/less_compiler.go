package web_resources

import "github.com/urfave/cli"

var LessCompilerCMD = cli.Command{
	Name:     "less_compiler",
	Aliases:  []string{"less_c"},
	Usage:    "Compiles a LESS file into a CSS file. Full support for all LESS features except @import.",
	Category: "网页资源",
	Action: func(c *cli.Context) error {
		return nil
	},
}
