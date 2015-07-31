package doublebuffer

import (
	"mygameengine/image"
)

type DoubleBuffer struct {
	buf1    *image.Image
	buf2    *image.Image
	current int8
}

func (d *DoubleBuffer) GetCurrentImage() *image.Image {
	if d.current == 1 {
		return d.buf1
	}
	return d.buf2
}

func (d *DoubleBuffer) GetPreviousImage() *image.Image {
	if d.current == 1 {
		return d.buf2
	}
	return d.buf1
}

func (d *DoubleBuffer) SwapBuffers() {
	d.current = d.current * -1
}

func New(width uint, height uint) *DoubleBuffer {
	doublebuffer := new(DoubleBuffer)
	doublebuffer.current = 1
	doublebuffer.buf1 = image.New(width, height)
	doublebuffer.buf2 = image.New(width, height)
	return doublebuffer
}
