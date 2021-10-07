package vpad

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type JustPressedButton struct {
	baseImg     *ebiten.Image
	normalOp    *ebiten.DrawImageOptions
	selectedOp  *ebiten.DrawImageOptions
	rectangle   image.Rectangle
	isSelected  bool
	isTriggered bool
	checkP      image.Point
}

func (b *JustPressedButton) SetLocation(x int, y int) {
	w, h := b.baseImg.Size()
	b.rectangle = image.Rect(x, y, x+w, y+h)

	b.normalOp.GeoM.Translate(float64(x), float64(y))
	b.selectedOp.GeoM.Concat(b.normalOp.GeoM)
	b.checkP = image.Point{}
}

func (b *JustPressedButton) Update() {
	b.updateSelect()
	b.updateTrigger()
}

func (b *JustPressedButton) updateSelect() {
	b.isSelected = false

	for _, tID := range ebiten.TouchIDs() {
		b.checkP.X, b.checkP.Y = ebiten.TouchPosition(tID)
		if b.checkP.In(b.rectangle) {
			b.isSelected = true
			return
		}
	}
}

func (b *JustPressedButton) updateTrigger() {
	b.isTriggered = false

	// JustPressed!
	for _, tID := range inpututil.JustPressedTouchIDs() {
		b.checkP.X, b.checkP.Y = ebiten.TouchPosition(tID)
		if b.checkP.In(b.rectangle) {
			b.isTriggered = true
			return
		}
	}
}

func (b *JustPressedButton) IsTriggered() bool {
	return b.isTriggered
}

func (b *JustPressedButton) Draw(screen *ebiten.Image) {
	if b.isSelected {
		screen.DrawImage(b.baseImg, b.selectedOp)
	} else {
		screen.DrawImage(b.baseImg, b.normalOp)
	}
}
