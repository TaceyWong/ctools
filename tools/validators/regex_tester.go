package validators

import "github.com/urfave/cli/v2"

var RegexTestCMD = cli.Command{
	Name:     "regex_test",
	Aliases:  []string{"rt"},
	Usage:    "Highlight every match in the original string so that you know exactly where a match occurs.",
	Category: "Validators",
	Action: func(c *cli.Context) error {
		return nil
	},
}
