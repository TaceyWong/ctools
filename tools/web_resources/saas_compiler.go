package web_resources

import "github.com/urfave/cli"

var SassCompilerCMD = cli.Command{
	Name:     "sass_compiler",
	Aliases:  []string{"sass_c"},
	Usage:    "Compiles a SASS file into a CSS file. Full support for all LESS features except @import.",
	Category: "Web Resources",
	Action: func(c *cli.Context) error {
		return nil
	},
}
