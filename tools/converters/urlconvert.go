package converters

import "github.com/urfave/cli/v2"

// 原始地址支持迅雷,快车,旋风的下载地址互转.

var URL2BTCMD = cli.Command{
	Name:     "url2bt",
	Aliases:  []string{"c2x"},
	Usage:    "convert a url-str into bt-str(xunlei-kuaiche-xuanfeng)",
	Category: "Converters",
	Action: func(c *cli.Context) error {

		return nil
	},
}
