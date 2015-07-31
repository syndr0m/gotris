package main

import (
	"fmt"
	"mygameengine"
	"mygameengine/image"
)

func main() {
	var screenWidth uint = 640
	var screenHeight uint = 480
	var fps uint = 25

	skylon, _ := image.Png("skylon.png")

	engine := mygameengine.New(screenWidth, screenHeight, fps)

	intro := engine.NewBoard()
	intro.OnKeyDown(func(key int) { fmt.Println("main: KEY DOWN") })
	intro.OnRepaint(func() {
		frame := intro.GetFrame()
		if frame < 255 {
			alpha := uint8(frame)
			screen := engine.GetScreenImage()
			screen.Blit(skylon)
			screen.DrawMask(0, 0, screenWidth, screenHeight, alpha)
		}
	})
	engine.SetCurrentBoard(intro)
	engine.Run()
}
