package validators

import "github.com/urfave/cli/v2"

var HTMLValidCMD = cli.Command{
	Name:     "html_valid",
	Aliases:  []string{"hv"},
	Usage:    "Validates the HTML string/file for well-formedness and compliance with w3c standards",
	Category: "Validators",
	Action: func(c *cli.Context) error {
		return nil
	},
}
