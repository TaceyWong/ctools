package web_resources

import "github.com/urfave/cli/v2"

var LessCompilerCMD = cli.Command{
	Name:     "less_compiler",
	Aliases:  []string{"less_c"},
	Usage:    "Compiles a LESS file into a CSS file. Full support for all LESS features except @import.",
	Category: "Web Resources",
	Action: func(c *cli.Context) error {
		return nil
	},
}
