package encoders_decoders

import "github.com/urfave/cli"

var Base64CMD = cli.Command{
	Name:     "base64",
	Aliases:  []string{"bs64"},
	Usage:    "Encodes or decodes a string so that it conforms to the Base64 Data Encodings specification (RFC 4648).",
	Category: "编码解码",
	Action: func(c *cli.Context) error {
		return nil
	},
}
