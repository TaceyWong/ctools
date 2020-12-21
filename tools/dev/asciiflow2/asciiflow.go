package asciiflow2

import (
	"fmt"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/TaceyWong/ctools/pkg/open"
	"github.com/urfave/cli/v2"
)

/*启动静态Web服务器
https://github.com/m3ng9i/ran
*/

// AsciiFlowCMD asciiflow cmd
var AsciiFlowCMD = cli.Command{
	Name:     "asciiflow",
	Aliases:  []string{"af"},
	Usage:    "tool of drawing ascii flow.",
	Category: "开发工具",
	Action: func(c *cli.Context) error {
		asciiflow()
		return nil
	},
}

func asciiflow() {
	go open.Start("http://127.0.0.1:3000/asciiflow")
	fmt.Println("Pls open http://127.0.0.1:3000/asciiflow in browser")
	// https://andrewbrookins.com/tech/golang-get-directory-of-the-current-file/
	// _, filename, _, _ := runtime.Caller(1)
	// fmt.Println(path.Join(path.Dir(filename), "static"))
	box := rice.MustFindBox("static")
	cssFileServer := http.StripPrefix("/asciiflow/", http.FileServer(box.HTTPBox()))
	http.Handle("/asciiflow/", cssFileServer)
	http.ListenAndServe(":3000", nil)
}
