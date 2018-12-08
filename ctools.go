package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/TaceyWong/ctools/tools"
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
	app.Copyright = "(c) 2018 - Forever N/A"
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
	app.Commands = append(app.Commands, tools.HtpasswdCMD)
	app.Commands = append(app.Commands, tools.IMailCMD)
	app.Commands = append(app.Commands, tools.HTTPCodeCMD)
	app.Commands = append(app.Commands, tools.RequestCMD)
	sort.Sort(cli.CommandsByName(app.Commands))
	sort.Sort(cli.FlagsByName(app.Flags))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
