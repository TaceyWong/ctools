# ctools

Coder👨‍💻 Tools🛠

**Now：Just a Test Repo**


## Desc

run :
```shell
$ctools
```

output:

```text

  .,-:::::::::::::::::   ...         ...      :::     .::::::.
  ,;;;'''';;;;;;;;''''.;;;;;;;.   .;;;;;;;.   ;;;    ;;;'    '
[[[            [[    ,[[     \[[,,[[     \[[, [[[    '[==/[[[[,
$$$            $$    $$$,     $$$$$$,     $$$ $$'      '''    $
'88bo,__,o,    88,   "888,_ _,88P"888,_ _,88Po88oo,.__88b    dP
  "YUMMMMMP"   MMM     "YMMMMMP"   "YMMMMMP" """"YUMMM "YMmMY"

Please type `ctools -h/--help` for the help of usage
```
run `ctools -h`:

output:
```text
NAME:
   ctools - Coder Toolbox,Programmer General Tools

USAGE:
   ctools [global options] command [command options] [arguments...]

VERSION:
   0.0.1dev

AUTHORS:
   Tacey Wong <xinyong.wang@qq.com>
   All Contributors

COMMANDS:
     htpasswd, hp  gen htpassword
     ...
     help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE  load configuration from FILE
   --lang value, -l value  language for the app (default: "english")
   --help, -h              show help
   --version, -v           print the version

COPYRIGHT:
   (c) 2018 - Forever N/A
```

Tree list of tools in **ctools**:




## Contribute

**ctools** uses these following third-party toolkits mainly：

+ [cli](https://github.com/urfave/cli) for parsing commandline args
+ [viper](https://github.com/spf13/viper) for loading configuration
+ [gocui](https://github.com/jroimartin/gocui) for Console UI (some tools maybe use it)
+ [go-prompt](https://github.com/c-bata/go-prompt) for interactive prompts
+ [survey](https://github.com/AlecAivazis/survey) for interactive prompts


code template:

```golang
package tools

/*
description of tool
*/

import (
    "github.com/urfave/cli"
    //what import
)

#ToolCMD desc
var ToolCMD = cli.Command{
	Name:    "tool-name",
	Aliases: []string{"tool-name-alias"},
	Usage:   "the function of tool",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
        },
        //flag
	},
	Action: Tool,//main function of tool
}

# Tool desc
func Tool (c *cli.Context) error {
		return nil
	}
```
place it into `ctools/tools/` and append the `ToolCMD`  to `app.Commands` in `ctools.go` like following:

```golang
app.Commands = append(app.Commands, tools.ToolCMD)
```
