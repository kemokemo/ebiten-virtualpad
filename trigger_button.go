package vpad

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// TriggerType is the type for the TriggerButton.
type TriggerType int

const (
	// JustRelease is the button to be triggered when just released only.
	JustRelease TriggerType = iota
	// Pressing is the button to be triggered during pressed every frame
	Pressing
)

type TriggerButton interface {
	SetLocation(x, y int)
	Update()
	IsTriggered() bool
	Draw(*ebiten.Image)
}

// NewTriggerButton returns a new TriggerButton.
func NewTriggerButton(img *ebiten.Image, tt TriggerType) TriggerButton {
	sop := &ebiten.DrawImageOptions{}
	sop.ColorM.Scale(colorScale(color.RGBA{0, 148, 255, 255}))
	switch tt {
	case JustRelease:
		return &JustReleaseButton{
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
	default:
		return nil
	}
}
