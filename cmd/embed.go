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
	//go:embed images/c_button.png
	c_button_png []byte
	//go:embed images/L_button.png
	L_button_png []byte
	//go:embed images/R_button.png
	R_button_png []byte
	//go:embed images/L_button_selected.png
	L_button_selected_png []byte
	//go:embed images/R_button_selected.png
	R_button_selected_png []byte

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
	c_button           *ebiten.Image
	L_button           *ebiten.Image
	R_button           *ebiten.Image
	L_button_selected  *ebiten.Image
	R_button_selected  *ebiten.Image
	directional_button *ebiten.Image
	directional_pad    *ebiten.Image
)

func init() {
	var err error

	a_button, err = loadSingleImage(a_button_png)
	if err != nil {
		log.Println("failed to load image: ", err)
	}
	b_button, err = loadSingleImage(b_button_png)
	if err != nil {
		log.Println("failed to load image: ", err)
	}
	c_button, err = loadSingleImage(c_button_png)
	if err != nil {
		log.Println("failed to load image: ", err)
	}
	L_button, err = loadSingleImage(L_button_png)
	if err != nil {
		log.Println("failed to load image: ", err)
	}
	R_button, err = loadSingleImage(R_button_png)
	if err != nil {
		log.Println("failed to load image: ", err)
	}
	L_button_selected, err = loadSingleImage(L_button_selected_png)
	if err != nil {
		log.Println("failed to load image: ", err)
	}
	R_button_selected, err = loadSingleImage(R_button_selected_png)
	if err != nil {
		log.Println("failed to load image: ", err)
	}

	directional_button, err = loadSingleImage(directional_button_png)
	if err != nil {
		log.Println("failed to load image: ", err)
	}
	directional_pad, err = loadSingleImage(directional_pad_png)
	if err != nil {
		log.Println("failed to load image: ", err)
	}
}
