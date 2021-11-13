package vpad

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type justPressedButton struct {
	baseImg     *ebiten.Image
	normalOp    *ebiten.DrawImageOptions
	selectedOp  *ebiten.DrawImageOptions
	rectangle   image.Rectangle
	isSelected  bool
	isTriggered bool
	checkP      image.Point
	cursP       image.Point
}

func (b *justPressedButton) SetLocation(x int, y int) {
	w, h := b.baseImg.Size()
	b.rectangle = image.Rect(x, y, x+w, y+h)

	b.normalOp.GeoM.Translate(float64(x), float64(y))
	b.selectedOp.GeoM.Concat(b.normalOp.GeoM)
	b.checkP = image.Point{}
	b.cursP = image.Point{}
}

func (b *justPressedButton) Update() {
	b.updateSelect()
	b.updateTrigger()
}

func (b *justPressedButton) updateSelect() {
	b.isSelected = false

	for _, tID := range ebiten.TouchIDs() {
		b.checkP.X, b.checkP.Y = ebiten.TouchPosition(tID)
		if b.checkP.In(b.rectangle) {
			b.isSelected = true
			return
		}
	}

	b.cursP.X, b.cursP.Y = ebiten.CursorPosition()
	if b.cursP.In(b.rectangle) && inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) > 0 {
		b.isSelected = true
		return
	}
}

func (b *justPressedButton) updateTrigger() {
	b.isTriggered = false

	// JustPressed!
	for _, tID := range inpututil.JustPressedTouchIDs() {
		b.checkP.X, b.checkP.Y = ebiten.TouchPosition(tID)
		if b.checkP.In(b.rectangle) {
			b.isTriggered = true
			return
		}
	}

	b.cursP.X, b.cursP.Y = ebiten.CursorPosition()
	if b.cursP.In(b.rectangle) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		b.isTriggered = true
		return
	}
}

func (b *justPressedButton) IsTriggered() bool {
	return b.isTriggered
}

func (b *justPressedButton) Draw(screen *ebiten.Image) {
	if b.isSelected {
		screen.DrawImage(b.baseImg, b.selectedOp)
	} else {
		screen.DrawImage(b.baseImg, b.normalOp)
	}
}
