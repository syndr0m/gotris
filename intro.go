package main

import (
	"fmt"
	"mygameengine"
	"mygameengine/image"
)

func NewBoardIntro(engine *mygameengine.MyGameEngine) *mygameengine.Board {
	skylon, _ := image.Png("skylon.png")

	intro := engine.NewBoard()
	intro.OnKeyDown(func(key int) { fmt.Println("main: KEY DOWN") })
	intro.OnRepaint(func(screen *image.Image) {
		frame := intro.GetFrame()
		if frame < 255 {
			alpha := uint8(frame)
			screen.Blit(skylon)
			screen.DrawMask(0, 0, engine.GetScreen().GetWidth(), engine.GetScreen().GetHeight(), alpha)
		}
	})
	return intro
}
