/*
 * @Description: 二维码生成
 * @Author: WangXinyong/TaceyWong
 * @Date: 2020-05-28 14:13:30
 * @LastEditors: WangXinyong/TaceyWong
 * @LastEditTime: 2020-05-28 14:27:23
 * @FilePath: /ctools/tools/fun/genqrcode.go
 */ 
package fun


import (
	"github.com/urfave/cli/v2"
	qrcode "github.com/skip2/go-qrcode"
)
//https://github.com/boombuler/barcode
//github.com/tuotoo/qrcode

// WhoisCMD Whois cli command
var QRCMD = cli.Command{
	Name:    "qrcode",
	Aliases: []string{"qr"},
	Usage:   "生成二维码",
	Category: "tools for fun",
	Action: QRCode,
}

func QRCode(c *cli.Context)error{
	//https://qrcode.51240.com/
	err := qrcode.WriteFile("https://example.org", qrcode.Medium, 256, "qr.png")
	return err
}