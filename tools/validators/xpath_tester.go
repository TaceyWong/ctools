package validators

import "github.com/urfave/cli/v2"

var XPathTestCMD = cli.Command{
	Name:     "xpath_test",
	Aliases:  []string{"xt"},
	Usage:    "Executes an XPath query against an XML file",
	Category: "Validators",
	Action: func(c *cli.Context) error {
		return nil
	},
}
