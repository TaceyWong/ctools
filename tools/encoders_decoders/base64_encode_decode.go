package encoders_decoders

import "github.com/urfave/cli/v2"

var Base64CMD = cli.Command{
	Name:     "base64",
	Aliases:  []string{"bs64"},
	Usage:    "Encodes or decodes a string so that it conforms to the Base64 Data Encodings specification (RFC 4648).",
	Category: "Encoders & Decoders",
	Action: func(c *cli.Context) error {
		return nil
	},
}
