package vpad

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type justPressedSelectButton struct {
	baseImg    *ebiten.Image
	normalOp   *ebiten.DrawImageOptions
	selectedOp *ebiten.DrawImageOptions
	rectangle  image.Rectangle
	isSelected bool
	checkP     image.Point
	cursP      image.Point
}

func (b *justPressedSelectButton) SetLocation(x int, y int) {
	w, h := b.baseImg.Size()
	b.rectangle = image.Rect(x, y, x+w, y+h)

	b.normalOp.GeoM.Translate(float64(x), float64(y))
	b.selectedOp.GeoM.Concat(b.normalOp.GeoM)
	b.checkP = image.Point{}
	b.cursP = image.Point{}
}

func (b *justPressedSelectButton) Update() {
	b.updateSelect()
}

func (b *justPressedSelectButton) updateSelect() {
	for _, tID := range inpututil.JustPressedTouchIDs() {
		b.checkP.X, b.checkP.Y = ebiten.TouchPosition(tID)
		if b.checkP.In(b.rectangle) {
			b.isSelected = !b.isSelected
			return
		}
	}

	b.cursP.X, b.cursP.Y = ebiten.CursorPosition()
	if b.cursP.In(b.rectangle) && inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) > 0 {
		b.isSelected = !b.isSelected
		return
	}
}

func (b *justPressedSelectButton) IsSelected() bool {
	return b.isSelected
}

func (b *justPressedSelectButton) SetSelectState(selected bool) {
	b.isSelected = selected
}

func (b *justPressedSelectButton) Draw(screen *ebiten.Image) {
	if b.isSelected {
		screen.DrawImage(b.baseImg, b.selectedOp)
	} else {
		screen.DrawImage(b.baseImg, b.normalOp)
	}
}
