package converters

import "github.com/urfave/cli"

// CSV2JSONCMD csv转json命令
var CSV2JSONCMD = cli.Command{
	Name:     "csv2json",
	Aliases:  []string{"c2j"},
	Usage:    "CSV转换为JSON",
	Category: "转换器",
	Action: func(c *cli.Context) error {
		return nil
	},
}
