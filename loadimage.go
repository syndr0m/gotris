package main

import (
	"image"
	_ "image/png"
	"os"
)

// FIXME: retourner la source
// FIXME: essayer de déplacer l'image !
// W00t
func loadImage(filename string) (image.Image, error) {
	infile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer infile.Close()

	src, _, err := image.Decode(infile)
	if err != nil {
		return nil, err
	}
	return src, nil
}
