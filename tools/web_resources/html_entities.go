package web_resources

import "github.com/urfave/cli/v2"

var HTMLEntitiesCMD = cli.Command{
	Name:     "html_entities",
	Aliases:  []string{"hentity"},
	Usage:    "HTML实体完整列表(包含数字和名称)",
	Category: "网页资源",
	Action: func(c *cli.Context) error {
		return nil
	},
}
