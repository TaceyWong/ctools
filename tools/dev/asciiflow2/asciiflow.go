package asciiflow2

import (
	"net/http"
	"path"
	"runtime"

	rice "github.com/GeertJohan/go.rice"
	"github.com/TaceyWong/ctools/pkg/open"
	"github.com/urfave/cli"
)

/*启动静态Web服务器
https://github.com/m3ng9i/ran
*/

var AsciiFlowCMD = cli.Command{
	Name:     "asciiflow",
	Aliases:  []string{"af"},
	Usage:    "tool of drawing ascii flow.",
	Category: "Dev-Tool",
	Action: func(c *cli.Context) error {
		asciiflow()
		return nil
	},
}

func asciiflow() {
	go open.Open("http://127.0.0.1:3000")
	// https://andrewbrookins.com/tech/golang-get-directory-of-the-current-file/
	_, filename, _, _ := runtime.Caller(1)
	box := rice.MustFindBox(path.Join(path.Dir(filename), "static"))
	cssFileServer := http.StripPrefix("/asciiflow/", http.FileServer(box.HTTPBox()))
	http.Handle("/asciiflow/", cssFileServer)
	http.ListenAndServe(":3000", nil)
}
