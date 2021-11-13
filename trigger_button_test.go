package vpad

import (
	"image/color"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestNewTriggerButton(t *testing.T) {
	img := ebiten.NewImage(1, 1)

	type args struct {
		img *ebiten.Image
		tt  TriggerType
		cl  color.RGBA
	}
	tests := []struct {
		name    string
		args    args
		isValid bool
	}{
		{"JustPressed", args{img: img, tt: JustPressed, cl: SelectColor}, true},
		{"JustReleased", args{img: img, tt: JustReleased, cl: SelectColor}, true},
		{"Pressing", args{img: img, tt: Pressing, cl: SelectColor}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTriggerButton(tt.args.img, tt.args.tt, tt.args.cl)
			if tt.isValid && got == nil {
				t.Errorf("NewTriggerButton() returns nil. :TriggerType = %v, isValid = %v", tt.args.tt, tt.isValid)
			}
		})
	}
}
