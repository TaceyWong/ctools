package encoders_decoders

import "github.com/urfave/cli/v2"

var URLEncodeDecodeCMD = cli.Command{
	Name:     "url_encode_decode",
	Aliases:  []string{"ued"},
	Usage:    "URL编码和解码(RFC 1738).",
	Category: "编码解码",
	Action: func(c *cli.Context) error {
		return nil
	},
}
