package vpad

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var SelectColor = color.RGBA{0, 148, 255, 255}

func colorScale(clr color.Color) (rf, gf, bf, af float64) {
	r, g, b, a := clr.RGBA()
	if a == 0 {
		return 0, 0, 0, 0
	}

	rf = float64(r) / float64(a)
	gf = float64(g) / float64(a)
	bf = float64(b) / float64(a)
	af = float64(a) / 0xffff
	return
}

func isTouched(touchedID ebiten.TouchID, bounds image.Rectangle) bool {
	x, y := ebiten.TouchPosition(touchedID)
	return image.Point{X: x, Y: y}.In(bounds)
}
