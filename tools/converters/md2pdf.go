package converters

import "github.com/urfave/cli/v2"

var MD2PDFCMD = cli.Command{
	Name:     "md2pdf",
	Usage:    "Markdown转换为PDF",
	Action:   md2PDF,
	Category: "转换工具",
}

func md2PDF(c *cli.Context) error {
	return nil
}
