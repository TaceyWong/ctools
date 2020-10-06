package converters

import "github.com/urfave/cli"

var XML2XSLCMD = cli.Command{
	Name:     "xml2xsl",
	Aliases:  []string{"x2xsl"},
	Usage:    "Transform an XML file using an XSL (EXtensible Stylesheet Language) file.",
	Category: "Converters",
	Action: func(c *cli.Context) error {
		return nil
	},
}
