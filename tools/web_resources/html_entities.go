package web_resources

import "github.com/urfave/cli"

var HTMLEntitiesCMD = cli.Command{
	Name:     "html_entities",
	Aliases:  []string{"hentity"},
	Usage:    "Complete list of HTML entities with their numbers and names",
	Category: "Web Resources",
	Action: func(c *cli.Context) error {
		return nil
	},
}
