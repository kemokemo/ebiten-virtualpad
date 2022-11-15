package vpad

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type SelectButton interface {
	SetLocation(x, y int)
	Update()
	IsSelected() bool
	Draw(screen *ebiten.Image)
	SetSelectState(selected bool)
	SetSelectKeys(keys []ebiten.Key)
	SetSelectImage(img *ebiten.Image)
}

// NewSelectButton returns a new SelectButton.
// Argument 'TriggerType' specifies the operation for which this button will be selected.
// Only 'JustReleased' and 'JustPressed' are available. If you specify others, this func returns nil.
func NewSelectButton(img *ebiten.Image, tt TriggerType, cl color.RGBA) SelectButton {
	sop := &ebiten.DrawImageOptions{}
	sop.ColorM.Scale(colorScale(cl))

	switch tt {
	case JustReleased:
		return &justReleasedSelectButton{
			baseImg:    img,
			normalOp:   &ebiten.DrawImageOptions{},
			selectedOp: sop,
			touches:    make(map[*touch]struct{}),
		}
	case JustPressed:
		return &justPressedSelectButton{
			baseImg:    img,
			normalOp:   &ebiten.DrawImageOptions{},
			selectedOp: sop,
		}
	default:
		return nil
	}
}
