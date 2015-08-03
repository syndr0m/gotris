package main

import (
	"fmt"
	"math"
	"mygameengine"
	"mygameengine/image"
)

func drawIntroBackground(screen *image.Image, frame uint64, imageIntroRotatingBg *image.Image) {
	var y int = int(math.Mod(float64(frame), 217))

	screen.BlitAt(imageIntroRotatingBg, 0, y-217)
	screen.BlitAt(imageIntroRotatingBg, 0, y)
	screen.BlitAt(imageIntroRotatingBg, 0, y+217)
	screen.BlitAt(imageIntroRotatingBg, 0, y+434)
}

const maskDuration = 2 // seconds

func drawIntroMask(screen *image.Image, frame uint64, fps uint) {
	var maxFrame uint64 = uint64(maskDuration * fps)

	if frame < maxFrame {
		black := mygameengine.COLOR_BLACK
		black.A = uint8(math.Max(0, 255-float64(frame)*(255/float64(maxFrame))))
		screen.DrawRectangle(0, 0, 640, 480, black)
	}
}

func drawIntroText(screen *image.Image, frame uint64, imageIntroText *image.Image) {
	var v int = int(math.Mod(float64(frame), 50))
	if v > 25 {
		screen.BlitAt(imageIntroText, 210, 370)
	}
}

func Intro(engine *mygameengine.MyGameEngine) *mygameengine.Board {
	engine.Assets().Png(IMG_ROTATING_BG)
	engine.Assets().Png(IMG_TITLE)
	engine.Assets().Png(IMG_PRESS_SPACE)

	intro := mygameengine.NewBoard()
	intro.OnKeyDown(func(key int) {
		fmt.Println("intro: KEY DOWN", key)
		if key == 1 {
			engine.Boards().SetCurrent(engine.Boards().Get("game"))
		}
	})
	intro.OnRepaint(func(screen *image.Image) {
		frame := intro.GetFrame()
		drawIntroBackground(screen, frame, engine.Assets().Get(IMG_ROTATING_BG))
		screen.BlitAt(engine.Assets().Get(IMG_TITLE), 130, 140)
		drawIntroText(screen, frame, engine.Assets().Get(IMG_PRESS_SPACE))
		drawIntroMask(screen, frame, engine.GetFps())
	})
	return intro
}

const IMG_ROTATING_BG string = "assets/images/intro-rotating-bg.png" // 640x217
const IMG_TITLE string = "assets/images/intro-title.png"             // 380x120
const IMG_PRESS_SPACE string = "assets/images/intro-press-space-to-start.png"
