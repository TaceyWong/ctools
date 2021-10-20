package converters

import "github.com/urfave/cli/v2"

var SplitPDFCMD = cli.Command{
	Name:     "splitpdf",
	Usage:    "切分PDF",
	Action:   splitPDF,
	Category: "转换工具",
}

func splitPDF(c *cli.Context) error {
	return nil
}
