package converters

import "github.com/urfave/cli/v2"

var MD2EpubCMD = cli.Command{
	Name:     "md2epub",
	Usage:    "Markdown转换为EPUB",
	Action:   md2Epub,
	Category: "转换工具",
}

func md2Epub(c *cli.Context) error {
	return nil
}
