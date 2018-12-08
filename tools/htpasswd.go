package tools

import (
	"fmt"

	"gopkg.in/AlecAivazis/survey.v1"
)

func init() {

}
// https://github.com/tg123/go-htpasswd
// https://github.com/ByteFlinger/htpasswd/blob/master/htpasswd.go

// HTPasswd Gen htpasswd
func HTPasswd() {
	// the questions to ask
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
				Help:    "The crypto type you want to apply",
			},
		},
		{
			Name: "password",
			Prompt: &survey.Password{
				Message: "Please type your password",
				Help:    "The password you want to used for this user",
			},
			
			Validate: survey.Required,
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
