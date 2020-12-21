package encoders_decoders

import "github.com/urfave/cli/v2"

var QRCodeCMD = cli.Command{
	Name:     "qrcode",
	Aliases:  []string{"qr"},
	Usage:    "二维码生成和识别",
	Category: "编码解码",
	Action: func(c *cli.Context) error {
		return nil
	},
}
