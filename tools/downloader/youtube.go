package downloader

import "github.com/urfave/cli/v2"

var YouTuBeCMD = cli.Command{
	Name:     "youtube",
	Usage:    "YouTuBe视频下载",
	Action:   youtube,
	Category: "下载器",
}

func youtube(c *cli.Context) error {
	return nil
}
