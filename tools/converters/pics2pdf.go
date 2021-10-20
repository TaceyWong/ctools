package converters

import "github.com/urfave/cli/v2"

var Pics2PDFCMD = cli.Command{
	Name:     "pics2pdf",
	Usage:    "图片转换为PDF",
	Action:   pics2pdf,
	Category: "转换工具",
}

func pics2pdf(c *cli.Context) error {
	return nil
}
