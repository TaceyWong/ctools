package string_utils

import "github.com/urfave/cli/v2"

var JavaDotNetEscapeCMD = cli.Command{
	Name:     "java_dotnet_escape",
	Aliases:  []string{"jde"},
	Usage:    "Escapes or unescapes a Java or .Net string removing traces of offending characters that could prevent compiling.",
	Category: "String Escaper & Utilities",
	Action: func(c *cli.Context) error {
		return nil
	},
}
