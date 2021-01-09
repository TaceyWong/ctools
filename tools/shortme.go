package tools

/*
short url tool
ori url -> short url & short url -> ori url
*/

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"gopkg.in/AlecAivazis/survey.v1"
	// surveyCore "gopkg.in/AlecAivazis/survey.v1/core"
)

// https://github.com/tg123/go-htpasswd
// https://github.com/ByteFlinger/htpasswd/blob/master/htpasswd.go

// HtpasswdCMD HTPasswd cli command
var ShortMeCMD = cli.Command{
	Name:    "shortme",
	Aliases: []string{"hp"},
	Usage:   "url -> short url",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "Language for this",
		},
		&cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	},
	Action: func(c *cli.Context) error {
		ShortMe()
		return nil
	},
}

// ShortMe encode & decode
func ShortMe() {
	// surveyCore.QuestionIcon = "😵"
	// surveyCore.HelpIcon = "x"
	var qs = []*survey.Question{
		{
			Name: "username",
			Prompt: &survey.Input{
				Message: "What is the user name?",
				Help:    "The user name you want to to use fot auth/login"},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "cryptotype",
			Prompt: &survey.Select{
				Message: "Choose a crypto type:",
				Options: []string{"Sha-1", "Md5", "Crypt"},
				Default: "Sha-1",
				Help:    "Md5:Apache only;Sha-1:Netscape and Apache ;Crypt:all Unix host",
			},
		},
		{
			Name: "password",
			Prompt: &survey.Password{
				Message: "Please type your password",
				Help:    "The password you want to used for this user",
			},

			Validate: survey.MinLength(6),
		},
	}
	result := struct {
		Name       string `survey:"username"`
		CryptoType string
		Password   string
	}{}

	// perform the questions
	err := survey.Ask(qs, &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s : %s : %s.\n", result.Name, result.CryptoType, result.Password)

}
