package validators

import "github.com/urfave/cli"

var JSONValidCMD = cli.Command{
	Name:     "json_valid",
	Aliases:  []string{"jv"},
	Usage:    "Validates a JSON string against RFC 4627",
	Category: "Validators",
	Action: func(c *cli.Context) error {
		return nil
	},
}
