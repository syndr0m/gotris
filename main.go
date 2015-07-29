package main

import (
	"fmt"
	"github.com/google/gxui"
	"mygameengine"
)

func main() {
	var screenWidth int = 640
	var screenHeight int = 480
	var fps int = 25

	skylon, _ := loadImage("skylon.png")

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
		fmt.Println("main: REPAINT")
		// 1er frame: on affiche l'image "skylon.png"
		var frame int = engine.GetFrame()
		if frame == 1 {
			engine.Blit(skylon)
		}
		if frame < screenWidth {
			for y := 0; y < screenHeight; y++ {
				engine.Plot(frame, y, mygameengine.COLOR_WHITE)
			}
		}
	})
	engine.Run()
}
