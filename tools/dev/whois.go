/*
 * @Description: WHOIS查询
 * @Author: WangXinyong/TaceyWong
 * @Date: 2020-05-27 14:28:18
 * @LastEditors: WangXinyong/TaceyWong
 * @LastEditTime: 2020-05-27 15:21:44
 * @FilePath: /ctools/tools/dev/whois.go
 */
package dev

import (
	"github.com/urfave/cli/v2"
	"github.com/levigross/grequests"
	"net/url"
	"fmt"
	"errors"
)
// WhoisCMD Whois cli command
var WhoisCMD = cli.Command{
	Name:    "whois",
	Aliases: []string{"ws"},
	Usage:   "查询Whois信息",
	Category: "tools for dev",
	Action:  Whois,
}

func Whois(c *cli.Context) error {
	if  c.NArg() < 1{
		return errors.New("请输入要查询的域名")
	}
	input := c.Args().Get(0)
	u, err := url.Parse(input)
    if err != nil {
        return err
	}

	// validUrl := "https://www.whois.net/AjaxService.svc/ValidateUrlExtension"
	hostname := u.Hostname()
	if hostname == ""{
		hostname = input
	}
	fmt.Println("正在查询"+hostname+"请稍候...")
	ro := &grequests.RequestOptions{
		Data: map[string]string{"domain_search": hostname,
			"domain_search_results": hostname},
		InsecureSkipVerify: true,
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36",
		
	}
	url := "https://www.whois.net/"
	resposne,err := grequests.Post(url,ro)
	fmt.Println(err,resposne.StatusCode)
	fmt.Println("whois工具还有待进一步开发")
	//需要进一步处理：校验域名；访问首页获取cookie类信息；提前valid；补足header；解析最后的结果
	return nil
}
