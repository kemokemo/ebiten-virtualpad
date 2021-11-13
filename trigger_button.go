package vpad

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type TriggerButton interface {
	SetLocation(x, y int)
	Update()
	IsTriggered() bool
	Draw(screen *ebiten.Image)
}

// NewTriggerButton returns a new TriggerButton.
func NewTriggerButton(img *ebiten.Image, tt TriggerType, cl color.RGBA) TriggerButton {
	sop := &ebiten.DrawImageOptions{}
	sop.ColorM.Scale(colorScale(cl))

	switch tt {
	case JustReleased:
		return &JustReleasedButton{
			baseImg:    img,
			normalOp:   &ebiten.DrawImageOptions{},
			selectedOp: sop,
			touches:    make(map[*touch]struct{}),
		}
	case Pressing:
		return &PressingButton{
			baseImg:    img,
			normalOp:   &ebiten.DrawImageOptions{},
			selectedOp: sop,
		}
	case JustPressed:
		return &JustPressedButton{
			baseImg:    img,
			normalOp:   &ebiten.DrawImageOptions{},
			selectedOp: sop,
		}
	default:
		return nil
	}
}
