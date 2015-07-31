package main

import (
	"fmt"
	"math"
	"mygameengine"
	"mygameengine/image"
)

func drawBackground(screen *image.Image, frame uint64, imageIntroRotatingBg *image.Image) {
	var y int = int(math.Mod(float64(frame), 217))

	screen.BlitAt(imageIntroRotatingBg, 0, y-217)
	screen.BlitAt(imageIntroRotatingBg, 0, y)
	screen.BlitAt(imageIntroRotatingBg, 0, y+217)
	screen.BlitAt(imageIntroRotatingBg, 0, y+434)
}

const maskDuration = 2 // seconds

func drawMask(screen *image.Image, frame uint64, fps uint) {
	var maxFrame uint64 = uint64(maskDuration * fps)

	if frame < maxFrame {
		black := mygameengine.COLOR_BLACK
		black.A = uint8(math.Max(0, 255-float64(frame)*(255/float64(maxFrame))))
		screen.DrawRectangle(0, 0, 640, 480, black)
	}
}

func NewBoardIntro(engine *mygameengine.MyGameEngine) *mygameengine.Board {
	imageIntroRotatingBg, _ := image.Png("assets/images/intro-rotating-bg.png") // 640x217
	imageIntroTitle, _ := image.Png("assets/images/intro-title.png")            // 380x120

	intro := engine.NewBoard()
	intro.OnKeyDown(func(key int) { fmt.Println("main: KEY DOWN") })
	intro.OnRepaint(func(screen *image.Image) {
		frame := intro.GetFrame()
		drawBackground(screen, frame, imageIntroRotatingBg)
		screen.BlitAt(imageIntroTitle, 130, 140)
		drawMask(screen, frame, engine.GetFps())
	})
	return intro
}
