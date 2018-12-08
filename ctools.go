package main

import (
	"fmt"
	"log"
	"os"

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
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		show()
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
