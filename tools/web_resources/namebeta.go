package web_resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

// NameBetaCMD https://github.com/TimothyYe/namebeta
var NameBetaCMD = cli.Command{
	Name:      "namebeta",
	Aliases:   []string{"nb"},
	UsageText: "ctools namebeta [options] DOMAIN-TO-QUERY...",
	Usage:     "command line domain query tool",
	Category:  "Web Resources",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "more, m",
			Usage: "Query more results with input domain",
		}, &cli.BoolFlag{
			Name:  "whois,w",
			Usage: "Query WHOIS information with input domain",
		},
	},
	Action: nameBetaAction,
}

const (
	domainURL = "https://namebeta.com/api/query"
	whoisURL  = "https://namebeta.com/api/whois"
	referURL  = "https://namebeta.com/search/%s"
	userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"
)

const (
	checkSymbol = "\u2714"
	crossSymbol = "\u2716"
)

func nameBetaAction(c *cli.Context) (err error) {
	if c.Args().Len() == 0 {
		cli.ShowCommandHelpAndExit(c, c.Command.Name, 0)
	}
	for _, domain := range c.Args().Slice() {
		fmt.Printf("Query [%s]'s ", color.HiBlueString(domain))
		if c.Bool("whois") {
			fmt.Println("whois")
			err = queryWhois(domain)
		} else {
			fmt.Println("Domain with more")
			err = queryDomain(domain, c.Bool("more"))
		}
		if err != nil {
			color.Red(fmt.Sprintf("%s %s \r\n", crossSymbol, err.Error()))
		}
	}

	return err
}

func queryWhois(domain string) error {

	//Init spinner
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Prefix = "Querying... "
	if err := s.Color("green"); err != nil {
		color.Red("Cannot set color")
		os.Exit(1)
	}
	s.Start()

	params := map[string]string{"domain": domain}
	result, err := getDomainInfo(whoisURL, domain, params)
	s.Stop()

	if err != nil {
		return err
	}

	if result[0].(bool) {
		fmt.Println()

		status := result[1].(map[string]interface{})["status"].(float64)

		if status == 1 {
			color.Red("NOT FOUND.")
		} else {
			color.Cyan(result[1].(map[string]interface{})["whois"].(string))
		}

		return nil
	}

	return fmt.Errorf("Failed to query domain: %s", domain)
}

func queryDomain(domain string, withMore bool) error {
	params := map[string]string{"q": domain}

	var resultMore []interface{}

	//Init spinner
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Prefix = "Querying... "
	if err := s.Color("green"); err != nil {
		color.Red("Failed to set color")
		os.Exit(1)
	}
	s.Start()

	result, err := getDomainInfo(domainURL, domain, params)

	// For more option, query special domains
	if withMore {
		params["special"] = "1"
		resultMore, err = getDomainInfo(domainURL, domain, params)

		if len(resultMore) > 0 && resultMore[0].(bool) {
			result[2] = append(result[2].([]interface{}), resultMore[2].([]interface{})...)
		}
	}

	s.Stop()

	if err != nil {
		return err
	}

	if result[0].(bool) {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Domain", "Available"})
		table.SetAlignment(tablewriter.ALIGN_CENTER)

		for _, v := range result[2].([]interface{}) {
			data := v.([]interface{})
			row := []string{data[0].(string)}

			switch data[1].(float64) {
			case 1:
				if runtime.GOOS == "windows" {
					row = append(row, checkSymbol)
				} else {
					row = append(row, color.GreenString(checkSymbol))
				}
			case 2:
				if runtime.GOOS == "windows" {
					row = append(row, crossSymbol)
				} else {
					row = append(row, color.GreenString(crossSymbol))
				}
			}

			table.Append(row)
		}
		table.Render()

		return nil
	}

	return fmt.Errorf("Failed to query domain: %s", domain)
}

func getDomainInfo(endpoint string, domain string, params map[string]string) ([]interface{}, error) {
	var result []interface{}

	form := url.Values{}
	for k, v := range params {
		form.Add(k, v)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Referer", fmt.Sprintf(referURL, domain))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("Failed to query domain: %s, endpoint: %s, error: %s", domain, endpoint, err.Error())
	}

	return result, nil
}
