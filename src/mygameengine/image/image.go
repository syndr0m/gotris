package image

/*
** RGBA image manipulation tool.
 */

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/png"
	"math"
	"os"
)

type Image struct {
	width  uint
	height uint
	buffer *image.RGBA
}

func (image *Image) GetWidth() uint         { return image.width }
func (image *Image) GetHeight() uint        { return image.height }
func (image *Image) GetBuffer() *image.RGBA { return image.buffer }

func (image *Image) Blit(img *Image) {
	image.BlitAt(img, 0, 0)
}

func (image *Image) BlitAt(img *Image, atx uint, aty uint) {
	var width uint = img.GetWidth()
	var height uint = img.GetHeight()

	var maxx uint = uint(math.Min(float64(image.width), float64(atx+width)))
	var maxy uint = uint(math.Min(float64(image.height), float64(aty+height)))

	for x := atx; x < maxx; x++ {
		for y := aty; y < maxy; y++ {
			image.Plot(x, y, img.At(x, y))
		}
	}
}

func (image *Image) At(x uint, y uint) color.RGBA {
	return image.buffer.RGBAAt(int(x), int(y))
}

func (image *Image) Plot(x uint, y uint, color color.RGBA) {
	image.buffer.SetRGBA(int(x), int(y), color)
}

func (image *Image) Mask(x uint, y uint, alpha uint8) {
	c := image.At(x, y)
	c.R = uint8(float64(c.R) * (255 - float64(alpha)) / 255)
	c.G = uint8(float64(c.G) * (255 - float64(alpha)) / 255)
	c.B = uint8(float64(c.B) * (255 - float64(alpha)) / 255)
	image.Plot(x, y, c)
}

func (image *Image) DrawRectangle(x1 uint, y1 uint, x2 uint, y2 uint, color color.RGBA) {
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			image.Plot(x, y, color)
		}
	}
}

func (image *Image) DrawMask(x1 uint, y1 uint, x2 uint, y2 uint, alpha uint8) {
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			image.Mask(x, y, alpha)
		}
	}
}

func New(width uint, height uint) *Image {
	newImage := new(Image)
	newImage.width = width
	newImage.height = height
	newImage.buffer = image.NewRGBA(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: int(width), Y: int(height)},
	})
	return newImage
}

func Png(filename string) (*Image, error) {
	infile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer infile.Close()

	src, _, err := image.Decode(infile)
	if err != nil {
		return nil, err
	}

	newImage := New(uint(src.Bounds().Max.X), uint(src.Bounds().Max.Y))
	draw.Draw(newImage.GetBuffer(), src.Bounds(), src, image.ZP, draw.Src)

	return newImage, nil
}
