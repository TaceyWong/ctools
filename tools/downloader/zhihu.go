package downloader

import "github.com/urfave/cli/v2"

var ZhihuCMD = cli.Command{
	Name:     "zhihu",
	Usage:    "知乎视频下载",
	Action:   zhihu,
	Category: "下载器",
}

func zhihu(c *cli.Context) error {
	return nil
}
