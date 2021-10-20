package converters

import "github.com/urfave/cli/v2"

var PathCMD = cli.Command{
	Name:     "path",
	Usage:    "路径格式转换",
	Action:   path,
	Category: "转换工具",
}

func path(c *cli.Context) error {
	return nil
}
