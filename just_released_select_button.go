package vpad

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// justReleasedSelectButton is implementation of the TriggerButton to be
// triggered when just released.
type justReleasedSelectButton struct {
	baseImg    *ebiten.Image
	normalOp   *ebiten.DrawImageOptions
	selectedOp *ebiten.DrawImageOptions
	rectangle  image.Rectangle
	isSelected bool
	touches    map[*touch]struct{}
	cursP      image.Point
}

// SetLocation sets the location to draw this button.
func (b *justReleasedSelectButton) SetLocation(x, y int) {
	w, h := b.baseImg.Size()
	b.rectangle = image.Rect(x, y, x+w, y+h)

	b.normalOp.GeoM.Translate(float64(x), float64(y))
	b.selectedOp.GeoM.Concat(b.normalOp.GeoM)
	b.cursP = image.Point{}
}

// Update updates the internal state of this button.
// Please call this before using IsTriggered method.
func (b *justReleasedSelectButton) Update() {
	b.updateSelect()
}

func (b *justReleasedSelectButton) updateSelect() {
	b.isSelected = false

	IDs := ebiten.TouchIDs()
	if len(IDs) >= 0 {
		for i := range IDs {
			if isTouched(IDs[i], b.rectangle) {
				b.isSelected = true
				return
			}
		}
	}

	b.cursP.X, b.cursP.Y = ebiten.CursorPosition()
	if b.cursP.In(b.rectangle) && inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) > 0 {
		b.isSelected = true
		return
	}
}

func (b *justReleasedSelectButton) IsSelected() bool {
	return b.isSelected
}

func (b *justReleasedSelectButton) SetSelectState(selected bool) {
	b.isSelected = selected
}

// Draw draws this button.
func (b *justReleasedSelectButton) Draw(screen *ebiten.Image) {
	if b.isSelected {
		screen.DrawImage(b.baseImg, b.selectedOp)
	} else {
		screen.DrawImage(b.baseImg, b.normalOp)
	}
}
