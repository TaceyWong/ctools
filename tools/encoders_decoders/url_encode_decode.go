package encoders_decoders

import "github.com/urfave/cli"

var URLEncodeDecodeCMD = cli.Command{
	Name:     "url_encode_decode",
	Aliases:  []string{"ued"},
	Usage:    "Encodes or decodes a string so that it conforms to the the Uniform Resource Locators Specification - URL (RFC 1738).",
	Category: "Encoders & Decoders",
	Action: func(c *cli.Context) error {
		return nil
	},
}
