package main

import (
	"mygameengine"
)

func main() {
	var screenWidth uint = 640
	var screenHeight uint = 480
	var fps uint = 25

	engine := mygameengine.New(screenWidth, screenHeight, fps)
	intro := NewBoardIntro(engine)
	engine.SetCurrentBoard(intro)
	engine.Run()
}
