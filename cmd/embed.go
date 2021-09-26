package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png" // to load png images
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed images/a_button.png
	a_button_png []byte

	//go:embed images/b_button.png
	b_button_png []byte

	//go:embed images/directional_button.png
	directional_button_png []byte

	//go:embed images/directional_pad.png
	directional_pad_png []byte
)

func loadSingleImage(b []byte) (*ebiten.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}

var (
	a_button           *ebiten.Image
	b_button           *ebiten.Image
	directional_button *ebiten.Image
	directional_pad    *ebiten.Image
)

func init() {
	var err error

	a_button, err = loadSingleImage(a_button_png)
	if err != nil {
		log.Println("failed to laod image: ", err)
	}

	b_button, err = loadSingleImage(b_button_png)
	if err != nil {
		log.Println("failed to laod image: ", err)
	}

	directional_button, err = loadSingleImage(directional_button_png)
	if err != nil {
		log.Println("failed to laod image: ", err)
	}

	directional_pad, err = loadSingleImage(directional_pad_png)
	if err != nil {
		log.Println("failed to laod image: ", err)
	}
}
