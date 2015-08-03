package main

import (
	"mygameengine"
)

func main() {
	var screenWidth uint = 640
	var screenHeight uint = 480
	var fps uint = 25

	engine := mygameengine.New(screenWidth, screenHeight, fps)
	engine.Boards().Register("intro", Intro(engine))
	engine.Boards().Register("game", Game(engine))
	engine.Boards().SetCurrent(engine.Boards().Get("intro"))
	engine.Run()
}
