package encoders_decoders

import "github.com/urfave/cli/v2"

var Base64CMD = cli.Command{
	Name:     "base64",
	Aliases:  []string{"bs64"},
	Usage:    "BASE64编码和解码(RFC 4648).",
	Category: "编码解码",
	Action: func(c *cli.Context) error {
		return nil
	},
}
