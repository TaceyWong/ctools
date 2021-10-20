package converters

import "github.com/urfave/cli/v2"

var PDF2PicsCMD = cli.Command{
	Name:     "pdf2pics",
	Usage:    "pdf转换为图片",
	Action:   pdf2Pics,
	Category: "转换工具",
}

func pdf2Pics(c *cli.Context) error {
	return nil
}
