package main

import (
	"fmt"
	"github.com/google/gxui"
	"mygameengine"
	"mygameengine/image"
)

func main() {
	var screenWidth uint = 640
	var screenHeight uint = 480
	var fps uint = 25

	skylon, _ := image.Png("skylon.png")

	engine := mygameengine.New(screenWidth, screenHeight, fps)
	engine.OnKeyDown(func(key int) {
		// fonction triggered a chaque appuis sur une touche
		fmt.Println("main: KEY DOWN")
		switch key {
		case int(gxui.KeyLeft):
			fmt.Println("main: left")
		case int(gxui.KeyRight):
			fmt.Println("main: right")
		case int(gxui.KeySpace):
			fmt.Println("main: space")
		}
	})
	engine.OnRepaint(func() {
		// fonction triggered à chaque frame
		// y placer le code qui redessine l'affichage
		var frame uint64 = engine.GetFrame()
		if frame < 255 {
			alpha := uint8(frame)
			screen := engine.GetScreenImage()
			screen.Blit(skylon)
			screen.DrawMask(0, 0, screenWidth, screenHeight, alpha)
		}
	})
	engine.Run()
}
