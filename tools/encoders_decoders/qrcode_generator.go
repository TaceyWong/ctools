package encoders_decoders

import "github.com/urfave/cli"

var QRCodeCMD = cli.Command{
	Name:     "qrcode",
	Aliases:  []string{"qr"},
	Usage:    "Generate QR Codes & Decode QR-Pic",
	Category: "Encoders & Decoders",
	Action: func(c *cli.Context) error {
		return nil
	},
}
