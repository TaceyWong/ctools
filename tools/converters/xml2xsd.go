package converters

import "github.com/urfave/cli/v2"

var XML2XSDCMD = cli.Command{
	Name:     "xml2xsd",
	Aliases:  []string{"x2x"},
	Usage:    "Generates a XSD (XML Schema) from a XML file.",
	Category: "Converters",
	Action: func(c *cli.Context) error {
		return nil
	},
}
