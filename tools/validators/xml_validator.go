package validators

import "github.com/urfave/cli"

var XMLValidCMD = cli.Command{
	Name:     "xml_valid",
	Aliases:  []string{"xv"},
	Usage:    "Validates the XML string/file against the specified XSD string/file",
	Category: "Validators",
	Action: func(c *cli.Context) error {
		return nil
	},
}
