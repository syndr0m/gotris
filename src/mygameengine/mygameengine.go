package mygameengine

import (
	"fmt"
	"mygameengine/doublebuffer"
	"mygameengine/image"
	"time"
)

type MyGameEngine struct {
	screen    *doublebuffer.DoubleBuffer
	fps       uint
	frame     uint64
	onKeyDown func(key int)
	onKeyUp   func(key int)
	onRepaint func()
	exit      chan int
}

type LoopFunc func(*MyGameEngine)

func (engine *MyGameEngine) OnKeyDown(f func(key int)) { engine.onKeyDown = f }
func (engine *MyGameEngine) OnKeyUp(f func(key int))   { engine.onKeyUp = f }
func (engine *MyGameEngine) OnRepaint(f func())        { engine.onRepaint = f }

func (engine *MyGameEngine) Run() {
	if engine.onRepaint == nil {
		panic("MyGameEngine.Run(): onRepaint function must exist")
	}

	// using gxui to open the window
	commands := make(chan Message)
	events := make(chan Message)
	go gxuiOpenWindow(engine.GetScreenImage().GetWidth(), engine.GetScreenImage().GetHeight(), engine.GetScreen(), commands, events)

	// engine-loop rendering
	ticker := time.NewTicker(time.Millisecond * time.Duration(1000/engine.fps))
	go func() {
		for {
			<-ticker.C
			// repaint
			engine.frame++
			// game repaint code
			engine.onRepaint()
			// switching buffer
			engine.GetScreen().SwapBuffers()
			// saving screen on the double buffer
			commands <- Message{MESSAGE_REPAINT, 0}
		}
	}()
	// engine-loop keys
	go func() {
		for {
			var m Message = <-events
			fmt.Println("EVENT MESSAGE RECEIVED")
			switch m.name {
			case MESSAGE_KEY_DOWN:
				fmt.Println("EVENT MESSAGE DISPATCHED TO KEYDOWN")
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

func (engine *MyGameEngine) GetFrame() uint64 {
	return engine.frame
}

func (engine *MyGameEngine) GetScreen() *doublebuffer.DoubleBuffer {
	return engine.screen
}

func (engine *MyGameEngine) GetScreenImage() *image.Image {
	return engine.screen.GetCurrentImage()
}

func New(screenWidth uint, screenHeight uint, fps uint) *MyGameEngine {
	engine := new(MyGameEngine)
	// default screen.
	engine.screen = doublebuffer.New(screenWidth, screenHeight)
	engine.fps = fps
	return engine
}
