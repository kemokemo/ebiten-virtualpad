package vpad

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type touch struct {
	id ebiten.TouchID
	x  int
	y  int
}

func (t *touch) Update() {
	if inpututil.IsTouchJustReleased(t.id) {
		return
	}
	t.x, t.y = ebiten.TouchPosition(t.id)
}

func (t *touch) IsReleased() bool {
	return inpututil.IsTouchJustReleased(t.id)
}
