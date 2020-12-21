/*
 * @Description:
 * @Author: WangXinyong/TaceyWong
 * @Date: 2020-05-27 14:49:56
 * @LastEditors: WangXinyong/TaceyWong
 * @LastEditTime: 2020-05-28 14:18:56
 * @FilePath: /ctools/ctools.go
 */
package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/TaceyWong/ctools/tools"
	"github.com/TaceyWong/ctools/tools/code_minifiers_beautifiers"
	"github.com/TaceyWong/ctools/tools/converters"
	"github.com/TaceyWong/ctools/tools/crypt_security"
	"github.com/TaceyWong/ctools/tools/encoders_decoders"
	"github.com/TaceyWong/ctools/tools/formatters"
	"github.com/TaceyWong/ctools/tools/fun"
	"github.com/TaceyWong/ctools/tools/string_utils"
	"github.com/TaceyWong/ctools/tools/validators"
	"github.com/TaceyWong/ctools/tools/web_resources"

	. "github.com/TaceyWong/ctools/tools/dev/asciiflow2"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

func show() {
	logo := `
  .,-:::::::::::::::::   ...         ...      :::     .::::::.
  ,;;;'''';;;;;;;;''''.;;;;;;;.   .;;;;;;;.   ;;;    ;;;'    '
[[[            [[    ,[[     \[[,,[[     \[[, [[[    '[==/[[[[,
$$$            $$    $$$,     $$$$$$,     $$$ $$'      '''    $
'88bo,__,o,    88,   "888,_ _,88P"888,_ _,88Po88oo,.__88b    dP
  "YUMMMMMP"   MMM     "YMMMMMP"   "YMMMMMP" """"YUMMM "YMmMY"
  `
	logo = `
   ______   ______            __    
  / ____/  /_  __/___  ____  / /____
 / /  ______/ / / __ \/ __ \/ / ___/
/ /__/_____/ / / /_/ / /_/ / (__  ) 
\____/    /_/  \____/\____/_/____/  
																	
	`
	fmt.Println(logo)
	fmt.Println("Please type `ctools -h/--help` for the help of usage")
}

func loadConfig() {
	appname := "ctools"
	viper.SetConfigName(appname)
	viper.AddConfigPath("/etc/" + appname + "/")
	viper.AddConfigPath("$HOME/." + appname)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}

func ctools() {
	app := cli.NewApp()
	app.Name = "ctools"
	app.Version = "0.0.1dev"
	app.Compiled = app.Compiled.UTC().Local()
	app.Authors = []*cli.Author{
		{
			Name:  "Tacey Wong",
			Email: "xinyong.wang@qq.com",
		},
		{
			Name: "所有贡献者",
		},
	}
	app.Copyright = "© 2018 - Forever Tacey Wong & 所有贡献者"
	app.Usage = "Coder Toolbox,Programmer General Tools"
	// app.UsageText = "contrive - demonstrating the available API"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "interactive",
			Aliases: []string{"i"},
			Usage:   "使用交互式模式",
		},
		&cli.StringFlag{
			Name:    "lang, l",
			Value:   "english",
			Usage:   "应用语言",
			EnvVars: []string{"CTOOLS_LANG"},
		},
		&cli.StringFlag{
			Name:  "config, c",
			Usage: "从`文件`加载配置",
		},
	}

	app.Commands = []*cli.Command{
		// Dev-Tools
		&AsciiFlowCMD,
		&tools.IMailCMD,
		&tools.RequestCMD,
		&tools.ShortMeCMD,

		&converters.JSON2GoCMD,
		&converters.JSON2XMLCMD,
		&converters.CSV2JSONCMD,
		&converters.CSV2XMLCMD,
		&converters.XML2JSONCMD,
		&converters.CURL2GoCMD,
		&converters.URL2BTCMD,
		&converters.XML2XSDCMD,
		&converters.Timestamp2DateCMD,
		&converters.XML2XSLCMD,

		&code_minifiers_beautifiers.CSSBeautifierCMD,
		&code_minifiers_beautifiers.CSSMinifierCMD,
		&code_minifiers_beautifiers.JavascriptBeautifierCMD,
		&code_minifiers_beautifiers.JavascriptMinifierCMD,

		&crypt_security.DigesterCMD,
		&crypt_security.HMACGeneratorCMD,
		&crypt_security.HtpasswdCMD,
		&crypt_security.MD5GeneratorCMD,
		&crypt_security.SHA256GeneratorCMD,
		&crypt_security.SHA512GeneratorCMD,
		&crypt_security.JWTRS256GeneratorCMD,

		&encoders_decoders.Base64CMD,
		&encoders_decoders.QRCodeCMD,
		&encoders_decoders.URLEncodeDecodeCMD,

		&formatters.HTMLFormatCMD,
		&formatters.JSONFormatCMD,
		&formatters.NginxConfigFormatCMD,
		&formatters.SQLFormatCMD,
		&formatters.XMLFormatCMD,

		&string_utils.StrUtilCMD,
		&string_utils.CSVEscapeCMD,
		&string_utils.HTMLEscapeCMD,
		&string_utils.JSONEscapeCMD,
		&string_utils.SQLEscapeCMD,
		&string_utils.XMLEscapeCMD,
		&string_utils.JavaDotNetEscapeCMD,

		&validators.CronGenCMD,
		&validators.HTMLValidCMD,
		&validators.JSONValidCMD,
		&validators.RegexTestCMD,
		&validators.XMLValidCMD,
		&validators.XPathTestCMD,

		&web_resources.CanadaProvinceCMD,
		&web_resources.ChinaProvinceCMD,
		&web_resources.HTMLEntitiesCMD,
		&web_resources.HTTPCodeCMD,
		&web_resources.I18nCMD,
		&web_resources.ISOCountryListCMD,
		&web_resources.LessCompilerCMD,
		&web_resources.LoremLpsumGeneratorCMD,
		&web_resources.MexcioStateCMD,
		&web_resources.MIMETypeListCMD,
		&web_resources.SassCompilerCMD,
		&web_resources.TimezoneListCMD,
		&web_resources.URLParserCMD,
		&web_resources.USAStateCMD,

		&fun.ConsolePicCMD,
	}
	app.HideHelpCommand = true
	sort.Sort(cli.CommandsByName(app.Commands))
	sort.Sort(cli.FlagsByName(app.Flags))
	app.Action = func(c *cli.Context) error {
		if c.Bool("i") {
			tools.InteractiveAction(c)
		} else {
			show()
		}
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// https://golanglibs.com
// https://www.freeformatter.com/
func main() {
	ctools()
}
