package converters

import "github.com/urfave/cli/v2"

// 原始地址支持迅雷,快车,旋风的下载地址互转.

var URL2BTCMD = cli.Command{
	Name:     "url2bt",
	Aliases:  []string{"c2x"},
	Usage:    "URL字符串转换为(迅雷-快车-旋风)地址",
	Category: "转换器",
	Action: func(c *cli.Context) error {

		return nil
	},
}
