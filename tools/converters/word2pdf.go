package converters

import "github.com/urfave/cli/v2"

var Word2PDFCMD = cli.Command{
	Name:     "word2pdf",
	Usage:    "Word转换为PDF",
	Action:   word2PDF,
	Category: "转换工具",
}

func word2PDF(c *cli.Context) error {
	return nil
}
