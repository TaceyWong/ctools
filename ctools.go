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
	"time"

	"github.com/TaceyWong/ctools/tools"
	"github.com/TaceyWong/ctools/tools/converters"
	"github.com/TaceyWong/ctools/tools/code_minifiers_beautifiers"
	"github.com/TaceyWong/ctools/tools/crypt_security"
	"github.com/TaceyWong/ctools/tools/encoders_decoders"
	"github.com/TaceyWong/ctools/tools/formatters"
	"github.com/TaceyWong/ctools/tools/string_utils"
	"github.com/TaceyWong/ctools/tools/validators"
	"github.com/TaceyWong/ctools/tools/web_resources"
	"github.com/TaceyWong/ctools/tools/fun"




	. "github.com/TaceyWong/ctools/tools/dev/asciiflow2"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
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
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Tacey Wong",
			Email: "xinyong.wang@qq.com",
		},
		cli.Author{
			Name: "All Contributors",
		},
	}
	app.Copyright = "© 2018 - Forever Tacey Wong and Contributors"
	app.Usage = "Coder Toolbox,Programmer General Tools"
	// app.UsageText = "contrive - demonstrating the available API"
	app.EnableBashCompletion = true
	app.Action = func(c *cli.Context) error {
		show()
		return nil
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "lang, l",
			Value:  "english",
			Usage:  "language for the app",
			EnvVar: "CTOOLS_LANG",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "load configuration from `FILE`",
		},
	}

	app.Commands = []cli.Command{}

	// Dev-Tools
	app.Commands = append(app.Commands, AsciiFlowCMD)
	app.Commands = append(app.Commands, tools.IMailCMD)
	app.Commands = append(app.Commands, tools.RequestCMD)
	app.Commands = append(app.Commands, tools.ShortMeCMD)


	// Converters
	app.Commands = append(app.Commands, converters.JSON2GoCMD)
	app.Commands = append(app.Commands, converters.JSON2XMLCMD)
	app.Commands = append(app.Commands, converters.CSV2JSONCMD)
	app.Commands = append(app.Commands, converters.CSV2XMLCMD)
	app.Commands = append(app.Commands, converters.XML2JSONCMD)
	app.Commands = append(app.Commands, converters.CURL2GoCMD)
	app.Commands = append(app.Commands, converters.URL2BTCMD)
	app.Commands = append(app.Commands, converters.XML2XSDCMD)
	app.Commands = append(app.Commands, converters.Timestamp2DateCMD)
	app.Commands = append(app.Commands, converters.XML2XSLCMD)

	// Code Minifiers&Beautifiers
	app.Commands = append(app.Commands, code_minifiers_beautifiers.CSSBeautifierCMD)
	app.Commands = append(app.Commands, code_minifiers_beautifiers.CSSMinifierCMD)
	app.Commands = append(app.Commands, code_minifiers_beautifiers.JavascriptBeautifierCMD)
	app.Commands = append(app.Commands, code_minifiers_beautifiers.JavascriptMinifierCMD)

    // Crypt & Security
	app.Commands = append(app.Commands, crypt_security.DigesterCMD)
	app.Commands = append(app.Commands, crypt_security.HMACGeneratorCMD)
	app.Commands = append(app.Commands, crypt_security.HtpasswdCMD)
	app.Commands = append(app.Commands, crypt_security.MD5GeneratorCMD)
	app.Commands = append(app.Commands, crypt_security.SHA256GeneratorCMD)
	app.Commands = append(app.Commands, crypt_security.SHA512GeneratorCMD)

    // Encoders & Decoders
	app.Commands = append(app.Commands, encoders_decoders.Base64CMD)
	app.Commands = append(app.Commands, encoders_decoders.QRCodeCMD)
	app.Commands = append(app.Commands, encoders_decoders.URLEncodeDecodeCMD)

    // Formatters
	app.Commands = append(app.Commands, formatters.HTMLFormatCMD)
	app.Commands = append(app.Commands, formatters.JSONFormatCMD)
	app.Commands = append(app.Commands, formatters.NginxConfigFormatCMD)
	app.Commands = append(app.Commands, formatters.SQLFormatCMD)
	app.Commands = append(app.Commands, formatters.XMLFormatCMD)

	// String Utils
	app.Commands = append(app.Commands, string_utils.StrUtilCMD)
	app.Commands = append(app.Commands, string_utils.CSVEscapeCMD)
	app.Commands = append(app.Commands, string_utils.HTMLEscapeCMD)
	app.Commands = append(app.Commands, string_utils.JSONEscapeCMD)
	app.Commands = append(app.Commands, string_utils.SQLEscapeCMD)
	app.Commands = append(app.Commands, string_utils.XMLEscapeCMD)
	app.Commands = append(app.Commands, string_utils.JavaDotNetEscapeCMD)

	// Validators
	app.Commands = append(app.Commands, validators.CronGenCMD)
	app.Commands = append(app.Commands, validators.HTMLValidCMD)
	app.Commands = append(app.Commands, validators.JSONValidCMD)
	app.Commands = append(app.Commands, validators.RegexTestCMD)
	app.Commands = append(app.Commands, validators.XMLValidCMD)
	app.Commands = append(app.Commands, validators.XPathTestCMD)


	// Web Resources
	app.Commands = append(app.Commands, web_resources.CanadaProvinceCMD)
	app.Commands = append(app.Commands, web_resources.ChinaProvinceCMD)
	app.Commands = append(app.Commands, web_resources.HTMLEntitiesCMD)
	app.Commands = append(app.Commands, web_resources.HTTPCodeCMD)
	app.Commands = append(app.Commands, web_resources.I18nCMD)
	app.Commands = append(app.Commands, web_resources.ISOCountryListCMD)
	app.Commands = append(app.Commands, web_resources.LessCompilerCMD )
	app.Commands = append(app.Commands, web_resources.LoremLpsumGeneratorCMD)
	app.Commands = append(app.Commands, web_resources.MexcioStateCMD)
	app.Commands = append(app.Commands, web_resources.MIMETypeListCMD)
	app.Commands = append(app.Commands, web_resources.SassCompilerCMD)
	app.Commands = append(app.Commands, web_resources.TimezoneListCMD)
	app.Commands = append(app.Commands, web_resources.URLParserCMD)
	app.Commands = append(app.Commands, web_resources.USAStateCMD)


    // Fun
	app.Commands = append(app.Commands, fun.ConsolePicCMD)



	sort.Sort(cli.CommandsByName(app.Commands))
	sort.Sort(cli.FlagsByName(app.Flags))
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
