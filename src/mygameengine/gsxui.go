package mygameengine

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
	"image"
)

func gxuiOpenWindow(width int, height int, buffer *image.RGBA, commands chan Message, events chan Message) {
	gl.StartDriver(func(driver gxui.Driver) {
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

		// repaint function
		go func() {
			for {
				<-commands
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
