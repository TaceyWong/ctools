/*
 * @Description: 填写描述
 * @Author: WangXinyong/TaceyWong
 * @Date: 2020-05-27 17:02:38
 * @LastEditors: WangXinyong/TaceyWong
 * @LastEditTime: 2020-05-27 17:02:53
 * @FilePath: /ctools/tools/fun/consolepic.go
 */ 
package fun


/* convert pixel to console-color-lable*/

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"strings"
    
    "github.com/urfave/cli/v2"
	"github.com/nfnt/resize"
)

const (
	end         = "\x1b[39;49m\n"
	pixelFormat = "\x1b[38;2;%d;%d;%d;48;2;%d;%d;%dm\u2580"
)

func convertImage(img image.Image) (output string) {
	bounds := img.Bounds()
	lines := make([]string, bounds.Max.Y)

	for row := 0; row < bounds.Max.Y; row += 2 {
		line := make([]string, bounds.Max.X+1)

		for col := 0; col <= bounds.Max.X; col++ {
			pixelUp := img.At(col, row)
			pixelDown := img.At(col, row+1)

			ru, gu, bu, _ := pixelUp.RGBA()
			rd, gd, bd, _ := pixelDown.RGBA()

			line[col] = fmt.Sprintf(pixelFormat,
				ru/257, gu/257, bu/257,
				rd/257, gd/257, bd/257)
		}
		line[bounds.Max.X] = end
		lines[row] = strings.Join(line, "")
	}
	return strings.Join(lines, "")
}

func resizeImage(img image.Image, cols int) image.Image {
	width := uint(cols)
	height := uint(float64(cols) / float64(img.Bounds().Max.X) * float64(img.Bounds().Max.Y) * 0.8)

	return resize.Resize(width, height, img, resize.Lanczos3)
}

func ConvertFromPath(path string, cols int) (output string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	return convertImage(resizeImage(img, cols)), nil
}

func ConvertFromURL(url string, cols int) (output string, err error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	img, _, err := image.Decode(response.Body)
	if err != nil {
		return "", err
	}
	return convertImage(resizeImage(img, cols)), nil
}


var ConsolePicCMD = cli.Command{
	Name:     "consolepic",
	Aliases:  []string{"console_p"},
	Usage:    "show a picture in console",
	Category: "Fun",
	Action: func(c *cli.Context) error {
		fmt.Println(ConvertFromURL(c.Args().Get(0), 128))
		return nil
	},
}
