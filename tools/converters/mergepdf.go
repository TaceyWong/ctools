package converters

import "github.com/urfave/cli/v2"

var MergePDFCMD = cli.Command{
	Name:     "mergepdf",
	Usage:    "合并PDF",
	Action:   mergePDF,
	Category: "转换工具",
}

func mergePDF(c *cli.Context) error {
	return nil
}
