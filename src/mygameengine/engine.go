package mygameengine

import (
	"fmt"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
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

const (
	MESSAGE_EXIT = iota
	MESSAGE_REPAINT
	MESSAGE_KEY_DOWN
	MESSAGE_KEY_UP
)

type Message struct {
	name  int
	value int
}

func gxuiOpenWindow(width int, height int, buffer *image.RGBA, commands chan Message, events chan Message) {
	fmt.Println("gxuiOpenWindow: startDriver")
	gl.StartDriver(func(driver gxui.Driver) {
		fmt.Println("gxuiOpenWindow: driver started")
		theme := flags.CreateTheme(driver)
		window := theme.CreateWindow(width, height, "MyGameEngine")
		window.SetScale(flags.DefaultScaleFactor)
		screen := theme.CreateImage()
		window.AddChild(screen)
		window.OnClose(func() {
			driver.Terminate()
			events <- Message{MESSAGE_EXIT, 0}
		})
		window.OnKeyDown(func(e gxui.KeyboardEvent) {
			events <- Message{MESSAGE_KEY_DOWN, int(e.Key)}
		})

		fmt.Println("gxuiOpenWindow: window opened")

		// repaint function
		go func() {
			for {
				fmt.Println("gxuiOpenWindow: wait for repaint message")
				<-commands
				fmt.Println("gxuiOpenWindow: REPAINT")
				last := screen.Texture()
				driver.CallSync(func() {
					texture := driver.CreateTexture(buffer, 1)
					screen.SetTexture(texture)
					if last != nil {
						last.Release()
					}
				})
			}
		}()
	})
}

type LoopFunc func(*MyGameEngine)

func (engine *MyGameEngine) OnKeyDown(f func(key int)) { engine.onKeyDown = f }
func (engine *MyGameEngine) OnKeyUp(f func(key int))   { engine.onKeyUp = f }
func (engine *MyGameEngine) OnRepaint(f func())        { engine.onRepaint = f }

func (engine *MyGameEngine) Run() {
	if engine.onRepaint == nil {
		panic("MyGameEngine.Run(): onRepaint function must exist")
	}

	fmt.Println("Run: create screenBuffer")
	engine.screenBuffer = image.NewRGBA(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: engine.screenWidth, Y: engine.screenHeight},
	})

	// using gxui to open the window
	fmt.Println("Run: create mygameengine <-> gsxui communication")
	commands := make(chan Message)
	events := make(chan Message)
	go gxuiOpenWindow(engine.screenWidth, engine.screenHeight, engine.screenBuffer, commands, events)

	// engine-loops
	fmt.Println("Run: create ticker")
	// rendering
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
	// keys
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

func (engine *MyGameEngine) GetFrame() int { return engine.frame }

func (engine *MyGameEngine) Blit(img image.Image) {
	draw.Draw(engine.screenBuffer, img.Bounds(), img, image.ZP, draw.Src)
}

func (engine *MyGameEngine) Plot(x int, y int, color color.RGBA) {
	engine.screenBuffer.SetRGBA(x, y, color)
}

var COLOR_WHITE color.RGBA = color.RGBA{R: 255, G: 255, B: 255, A: 255}
var COLOR_BLACK color.RGBA = color.RGBA{R: 0, G: 0, B: 0, A: 255}
var COLOR_RED color.RGBA = color.RGBA{R: 255, G: 0, B: 0, A: 255}
var COLOR_GREEN color.RGBA = color.RGBA{R: 0, G: 255, B: 0, A: 255}
var COLOR_BLUE color.RGBA = color.RGBA{R: 0, G: 0, B: 255, A: 255}

func New(screenWidth int, screenHeight int, fps int) *MyGameEngine {
	engine := new(MyGameEngine)
	engine.screenWidth = screenWidth
	engine.screenHeight = screenHeight
	engine.fps = fps
	return engine
}
