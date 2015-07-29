package mygameengine

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"time"
)

type MyGameEngine struct {
	screenWidth  int
	screenHeight int
	fps          int
	frame        int
	screenBuffer *image.RGBA
	onKeyDown    func(key int)
	onKeyUp      func(key int)
	onRepaint    func()
	exit         chan int
}

type LoopFunc func(*MyGameEngine)

func (engine *MyGameEngine) OnKeyDown(f func(key int)) { engine.onKeyDown = f }
func (engine *MyGameEngine) OnKeyUp(f func(key int))   { engine.onKeyUp = f }
func (engine *MyGameEngine) OnRepaint(f func())        { engine.onRepaint = f }

func (engine *MyGameEngine) Run() {
	if engine.onRepaint == nil {
		panic("MyGameEngine.Run(): onRepaint function must exist")
	}

	engine.screenBuffer = image.NewRGBA(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: engine.screenWidth, Y: engine.screenHeight},
	})

	// using gxui to open the window
	commands := make(chan Message)
	events := make(chan Message)
	go gxuiOpenWindow(engine.screenWidth, engine.screenHeight, engine.screenBuffer, commands, events)

	// engine-loop rendering
	ticker := time.NewTicker(time.Millisecond * time.Duration(1000/engine.fps))
	go func() {
		for {
			<-ticker.C
			fmt.Println("tick", engine.frame)
			// repaint
			engine.frame++
			engine.onRepaint()
			commands <- Message{MESSAGE_REPAINT, 0}
		}
	}()
	// engine-loop keys
	go func() {
		for {
			var m Message = <-events
			switch m.name {
			case MESSAGE_KEY_DOWN:
				engine.onKeyDown(m.value)
			case MESSAGE_EXIT:
				engine.Stop()
			}
		}
	}()

	engine.exit = make(chan int)
	<-engine.exit
}

func (engine *MyGameEngine) Stop() {
	engine.exit <- 42
}

func (engine *MyGameEngine) GetFrame() int {
	return engine.frame
}

func (engine *MyGameEngine) Blit(img image.Image) {
	draw.Draw(engine.screenBuffer, img.Bounds(), img, image.ZP, draw.Src)
}

func (engine *MyGameEngine) Plot(x int, y int, color color.RGBA) {
	engine.screenBuffer.SetRGBA(x, y, color)
}

func New(screenWidth int, screenHeight int, fps int) *MyGameEngine {
	engine := new(MyGameEngine)
	engine.screenWidth = screenWidth
	engine.screenHeight = screenHeight
	engine.fps = fps
	return engine
}
