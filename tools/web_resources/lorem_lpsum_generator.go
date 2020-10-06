package web_resources

import "github.com/urfave/cli"

var LoremLpsumGeneratorCMD = cli.Command{
	Name:     "lorem_lpsum_generator",
	Aliases:  []string{"llg"},
	Usage:    "choose how many sentences, paragraphs or list items you want. ",
	Category: "Web Resources",
	Action: func(c *cli.Context) error {
		return nil
	},
}
