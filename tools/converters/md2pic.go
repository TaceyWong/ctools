package converters

import "github.com/urfave/cli/v2"

var MD2PicCMD = cli.Command{
	Name:     "md2pic",
	Usage:    "Markdown转换为图片",
	Action:   md2Pic,
	Category: "转换工具",
}

func md2Pic(c *cli.Context) error {
	return nil
}
