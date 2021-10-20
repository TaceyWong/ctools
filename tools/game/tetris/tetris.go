package tetris

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/urfave/cli/v2"
)

var TetrisCMD = cli.Command{
	Name:   "tetris",
	Usage:  "小游戏俄罗斯方块",
	Category: "小游戏",
	Action: tetris,
}

var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")

func tetris(c *cli.Context) error {
	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	ebiten.SetWindowSize(ScreenWidth*2, ScreenHeight*2)
	ebiten.SetWindowTitle("俄罗斯方块")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
	return nil
}
