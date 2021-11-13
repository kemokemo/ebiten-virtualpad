package vpad

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// justReleasedButton is implementation of the TriggerButton to be
// triggered when just released.
type justReleasedButton struct {
	baseImg     *ebiten.Image
	normalOp    *ebiten.DrawImageOptions
	selectedOp  *ebiten.DrawImageOptions
	rectangle   image.Rectangle
	isSelected  bool
	isTriggered bool
	touches     map[*touch]struct{}
	cursP       image.Point
}

// SetLocation sets the location to draw this button.
func (b *justReleasedButton) SetLocation(x, y int) {
	w, h := b.baseImg.Size()
	b.rectangle = image.Rect(x, y, x+w, y+h)

	b.normalOp.GeoM.Translate(float64(x), float64(y))
	b.selectedOp.GeoM.Concat(b.normalOp.GeoM)
	b.cursP = image.Point{}
}

// Update updates the internal state of this button.
// Please call this before using IsTriggered method.
func (b *justReleasedButton) Update() {
	b.updateSelect()
	b.updateTrigger()
}

func (b *justReleasedButton) updateSelect() {
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

func (b *justReleasedButton) updateTrigger() {
	b.isTriggered = false

	IDs := inpututil.JustPressedTouchIDs()
	if len(IDs) != 0 {
		for _, id := range IDs {
			b.touches[&touch{id: id}] = struct{}{}
		}
	}

	for t := range b.touches {
		t.Update()
		if t.IsReleased() {
			delete(b.touches, t)
			in := image.Point{t.x, t.y}.In(b.rectangle)
			if in {
				b.isTriggered = true
				return
			}
		}
	}

	b.cursP.X, b.cursP.Y = ebiten.CursorPosition()
	if b.cursP.In(b.rectangle) && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		b.isTriggered = true
		return
	}
}

// IsTriggered returns the state of this trigger is pressed.
// If result is 'true', this is pressed now.
func (b *justReleasedButton) IsTriggered() bool {
	return b.isTriggered
}

// Draw draws this button.
func (b *justReleasedButton) Draw(screen *ebiten.Image) {
	if b.isSelected {
		screen.DrawImage(b.baseImg, b.selectedOp)
	} else {
		screen.DrawImage(b.baseImg, b.normalOp)
	}
}
