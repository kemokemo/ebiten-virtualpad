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
	keys       []ebiten.Key
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
				b.isSelected = !b.isSelected
			}
		}
	}

	b.cursP.X, b.cursP.Y = ebiten.CursorPosition()
	if b.cursP.In(b.rectangle) && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		b.isSelected = !b.isSelected
		return
	}

	for i := range b.keys {
		if inpututil.IsKeyJustReleased(b.keys[i]) {
			b.isSelected = !b.isSelected
			return
		}
	}
}

func (b *justReleasedSelectButton) IsSelected() bool {
	return b.isSelected
}

func (b *justReleasedSelectButton) SetSelectState(selected bool) {
	b.isSelected = selected
}

func (b *justReleasedSelectButton) SetSelectKeys(keys []ebiten.Key) {
	b.keys = keys
}

// Draw draws this button.
func (b *justReleasedSelectButton) Draw(screen *ebiten.Image) {
	if b.isSelected {
		screen.DrawImage(b.baseImg, b.selectedOp)
	} else {
		screen.DrawImage(b.baseImg, b.normalOp)
	}
}
