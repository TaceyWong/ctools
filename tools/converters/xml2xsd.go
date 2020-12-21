package converters

import "github.com/urfave/cli/v2"

var XML2XSDCMD = cli.Command{
	Name:     "xml2xsd",
	Aliases:  []string{"x2x"},
	Usage:    "XML转换为XSD",
	Category: "转换器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
