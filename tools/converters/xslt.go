package converters

import "github.com/urfave/cli/v2"

var XML2XSLCMD = cli.Command{
	Name:     "xml2xsl",
	Aliases:  []string{"x2xsl"},
	Usage:    "XML转换为XSL",
	Category: "转换器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
