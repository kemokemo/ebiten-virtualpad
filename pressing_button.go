package vpad

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// pressingButton is implementation of the TriggerButton to be
// triggered during being pressed.
type pressingButton struct {
	baseImg     *ebiten.Image
	normalOp    *ebiten.DrawImageOptions
	selectedOp  *ebiten.DrawImageOptions
	rectangle   image.Rectangle
	isSelected  bool
	isTriggered bool
	checkP      image.Point
	cursP       image.Point
}

// SetLocation sets the location to draw this button.
func (b *pressingButton) SetLocation(x, y int) {
	w, h := b.baseImg.Size()
	b.rectangle = image.Rect(x, y, x+w, y+h)

	b.normalOp.GeoM.Translate(float64(x), float64(y))
	b.selectedOp.GeoM.Concat(b.normalOp.GeoM)
	b.checkP = image.Point{}
	b.cursP = image.Point{}
}

// Update updates the internal state of this button.
// Please call this before using IsTriggered method.
func (b *pressingButton) Update() {
	b.updateSelect()
	b.updateTrigger()
}

func (b *pressingButton) updateSelect() {
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

func (b *pressingButton) updateTrigger() {
	b.isTriggered = false

	for _, tID := range ebiten.TouchIDs() {
		if inpututil.TouchPressDuration(tID) > 0 {
			b.checkP.X, b.checkP.Y = ebiten.TouchPosition(tID)
			if b.checkP.In(b.rectangle) {
				b.isTriggered = true
				return
			}
		}
	}

	b.cursP.X, b.cursP.Y = ebiten.CursorPosition()
	if b.cursP.In(b.rectangle) && inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) > 0 {
		b.isTriggered = true
		return
	}
}

// IsTriggered returns the state of this trigger is pressed.
// If result is 'true', this is pressed now.
func (b *pressingButton) IsTriggered() bool {
	return b.isTriggered
}

// Draw draws this button.
func (b *pressingButton) Draw(screen *ebiten.Image) {
	if b.isSelected {
		screen.DrawImage(b.baseImg, b.selectedOp)
	} else {
		screen.DrawImage(b.baseImg, b.normalOp)
	}
}
