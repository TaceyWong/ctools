package downloader

import "github.com/urfave/cli/v2"

var BilibiliCMD = cli.Command{
	Name:     "bilibili",
	Usage:    "哔哩哔哩视频下载",
	Action:   bilibili,
	Category: "下载器",
}

func bilibili(c *cli.Context) error {
	return nil
}
