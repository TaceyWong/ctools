package web_resources

import "github.com/urfave/cli"

var URLParserCMD = cli.Command{
	Name:     "url_parser",
	Aliases:  []string{"url_p"},
	Usage:    " parses a URL into its individual components (protocol, scheme, userinfo, host, port, authority, path, resource, etc.)",
	Category: "Web Resources",
	Action: func(c *cli.Context) error {
		return nil
	},
}
