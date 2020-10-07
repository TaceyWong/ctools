# ctools

[![Go Report Card](https://goreportcard.com/badge/github.com/TaceyWong/ctools)](https://goreportcard.com/report/github.com/TaceyWong/ctools)

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
   httpcode, hc  search http code info
   imail, im     send email
   request, rq   send http request,more friendly
   shortme, hp   url -> short url
   help, h       Shows a list of commands or help for one command

   Code Minifiers / Beautifier:
     css_beauty, css_p             Formats your CSS files for optimal readability.
     css_mini, css_m               Compresses a CSS string/file with no possible side-effect
     javascript_beauty, js_beauty  Formats your JS files for optimal readability
     javascript_mini, js_mini      Compresses a JavaScript string/file with no possible side-effect

   Converters:
     csv2json, c2j        convert a CSV file into an JSON file
     csv2xml, c2x         convert a CSV file into an XML file
     curl2go, curl2g       turns a curl command into Go code.
     json2go, c2g         convert a JSON file into an go struct file
     json2xml, c2x        convert a JSON file into an xml file
     timestamp2date, t2d  Converts an epoch/unix timestamp into a human readable date
     url2bt, c2x          convert a url-str into bt-str(xunlei-kuaiche-xuanfeng)
     xml2json, x2j        convert a XML file into an JSON file
     xml2xsd, x2x         Generates a XSD (XML Schema) from a XML file.
     xml2xsl, x2xsl       Transform an XML file using an XSL (EXtensible Stylesheet Language) file.

   Cryptography & Security:
     digester, digest      Computes a digest from a string using different algorithms.
     hmac_gen, hmac_g      Computes a Hash-based message authentication code (HMAC) using different algorithms.
     htpasswd, hp          gen httpassword
     md5_gen, md5_g        Computes a digest from a string using MD5
     sha256_gen, sha256_g  Computes a digest from a string using SHA-256.
     sha512_gen, sha512_g  Computes a digest from a string using SHA-512.

   Dev-Tool:
     asciiflow, af  tool of drawing ascii flow.

   Encoders & Decoders:
     base64, bs64            Encodes or decodes a string so that it conforms to the Base64 Data Encodings specification (RFC 4648).
     qrcode, qr              Generate QR Codes & Decode QR-Pic
     url_encode_decode, ued  Encodes or decodes a string so that it conforms to the the Uniform Resource Locators Specification - URL (RFC 1738).

   Formatters:
     html_format, hf                  Format HTML Text
     json_format, jf                  Format JSON Text
     nginx_conf_format, nginx_conf_f  Format HTML Text
     sql_format, sf                   Format SQL Text
     xml_format, hf                   Format XML Text

   Fun:
     consolepic, console_p  show a picture in console

   String Escaper & Utilities:
     csv_escape, ce           Escapes or unescapes a CSV string removing traces of offending characters that could prevent parsing.
     html_escape, he          Escapes or unescapes an HTML file removing traces of offending characters that could be wrongfully interpreted as markup.
     java_dotnet_escape, jde  Escapes or unescapes a Java or .Net string removing traces of offending characters that could prevent compiling.
     json_escape, jsone       Escapes or unescapes a JSON string removing traces of offending characters that could prevent parsing.
     sql_escape, sqle         Escapes or unescapes a SQL string removing traces of offending characters that could prevent execution.
     str_util, su             A couple of online string utilities written in JavaScript.
     xml_escape, xe           Escapes or unescapes an XML file removing traces of offending characters that could be wrongfully interpreted as markup.

   Validators:
     cron_gen, cg    Generate a quartz cron expression with an easy to use interface
     html_valid, hv  Validates the HTML string/file for well-formedness and compliance with w3c standards
     json_valid, jv  Validates a JSON string against RFC 4627
     regex_test, rt  Highlight every match in the original string so that you know exactly where a match occurs.
     xml_valid, xv   Validates the XML string/file against the specified XSD string/file
     xpath_test, xt  Executes an XPath query against an XML file

   Web Resources:
     canada_province, cprovince   snippets to generate a list of provinces and territories for Canada
     china_province, chinap       snippets to generate a list of provinces and territories for China
     html_entities, hentity      Complete list of HTML entities with their numbers and names
     i18n, i18                    list of formatting standards for dates, time, decimals, currencies and more with code snippets
     iso_country_list, icl        snippets to generate a list of countries using the ISO-3166-1 and ISO-3166-2 codes.
     less_compiler, less_c       Compiles a LESS file into a CSS file. Full support for all LESS features except @import.
     lorem_lpsum_generator, llg  choose how many sentences, paragraphs or list items you want.
     mexico_state, mexcio_s      snippets to generate a list state for Canada
     mime_type_list, mtl         Full list of MIME types
     sass_compiler, sass_c       Compiles a SASS file into a CSS file. Full support for all LESS features except @import.
     timezone_list, timezone_l   snippets to generate a list of time zones
     url_parser, url_p            parses a URL into its individual components (protocol, scheme, userinfo, host, port, authority, path, resource, etc.)
     usa_state, mexcio_s         Snippets to generate a list of states for the United States

GLOBAL OPTIONS:
   --config FILE, -c FILE  load configuration from FILE
   --lang value, -l value  language for the app (default: "english") [$CTOOLS_LANG]
   --help, -h              show help
   --version, -v           print the version

COPYRIGHT:
   © 2018 - Forever Tacey Wong and Contributors
```




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
