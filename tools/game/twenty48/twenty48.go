package twenty48

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/urfave/cli/v2"
)

var Twenty48CMD = cli.Command{
	Name:     "2048",
	Usage:    "小游戏2048",
	Category: "小游戏",
	Action:   twenty48,
}

func twenty48(c *cli.Context) error {
	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("2048")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
	return err
}
