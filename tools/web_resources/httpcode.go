package web_resources

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/urfave/cli"
)

var HTTPCodeCMD = cli.Command{
	Name:    "httpcode",
	Aliases: []string{"hc"},
	Usage:   "search http code info",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "Language for this",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	},
	Action: func(c *cli.Context) error {
		HTTPCode()
		return nil
	},
}


// HTTPCode x
func HTTPCode() {
	fmt.Println("Just type code and tab.")
	for true {
		t := prompt.Input("> ", suggester)
		fmt.Println("You selected " + t)
	}
}

func suggester(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

// StatusCode holds info http status code
type StatusCode struct {
	Code string
	Cate string
	ShortDesc string
	LongDesc string
}

func (c *StatusCode)String() string{
	return c.Code+c.Cate+c.ShortDesc+c.LongDesc
}

var codes = []*StatusCode{
	{
		Code: "100",
		Cate: "信息响应",
		ShortDesc: "Continue",
		LongDesc: `这个临时响应表明，迄今为止的所有内容都是可行的，客户端应该继续请求，如果已经完成，则忽略它。`,
	},
	{
		Code:"101",
		Cate:"信息响应",
		ShortDesc:"Switching Protocol",
		LongDesc:`该代码是响应客户端的 Upgrade 标头发送的，并且指示服务器也正在切换的协议。`,
	},
}


