package vpad

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// DirectionalPad is the directional pad for a game.
type DirectionalPad struct {
	baseImg           *ebiten.Image
	buttons           map[Direction]*directionalButton
	selectedDirection Direction
	op                *ebiten.DrawImageOptions
}

// NewDirectionalPad returns a new DirectionalPad.
func NewDirectionalPad(pad, button *ebiten.Image, cl color.RGBA) DirectionalPad {
	dp := DirectionalPad{
		baseImg: pad,
		op:      &ebiten.DrawImageOptions{},
	}
	dp.createButtons(button, cl)
	return dp
}

func (dp *DirectionalPad) createButtons(img *ebiten.Image, cl color.RGBA) {
	if dp.buttons == nil {
		dp.buttons = make(map[Direction]*directionalButton, 4)
	}

	ds := []Direction{Left, Upper, Right, Lower}
	for _, direc := range ds {
		b := newDirectionalButton(img, direc, cl)
		dp.buttons[direc] = b
	}
}

// SetLocation sets the location to draw this directional pad.
func (dp *DirectionalPad) SetLocation(x, y int) {
	dp.op.GeoM.Translate(float64(x), float64(y))

	wp, _ := dp.baseImg.Size()
	halfWp := int(wp / 2)
	wb, _ := dp.buttons[Left].Size()
	halfWb := int(wb / 2)
	outerMargin := int(0.25 * float64(halfWp))
	innerMargin := int(0.08 * float64(halfWp))

	dp.buttons[Left].SetLocation(x+outerMargin, y+halfWp-halfWb)
	dp.buttons[Upper].SetLocation(x+halfWp-halfWb, y+outerMargin)
	dp.buttons[Right].SetLocation(x+halfWp+innerMargin, y+halfWp-halfWb)
	dp.buttons[Lower].SetLocation(x+halfWp-halfWb, y+halfWp+innerMargin)
}

// Update updates the internal status of this struct.
func (dp *DirectionalPad) Update() {
	dp.updateDirection()
	dp.updateButtons()
}

func (dp *DirectionalPad) updateDirection() {
	IDs := ebiten.TouchIDs()

	if len(IDs) == 0 {
		// There are no touches
		dp.selectedDirection = None
		return
	}

	directions := []Direction{}
	// Find the newly touched direction key.
	for index := range IDs {
		for key := range dp.buttons {
			if isTouched(IDs[index], dp.buttons[key].GetRectangle()) {
				directions = append(directions, key)
			}
		}
	}
	dp.defineDirection(directions)
}

func (dp *DirectionalPad) defineDirection(directions []Direction) {
	current := None
	for index := range directions {
		current = getMergedDirection(current, directions[index])
	}
	dp.selectedDirection = current
}

func (dp *DirectionalPad) updateButtons() {
	for key := range dp.buttons {
		dp.buttons[key].SelectButton(false)
	}
	if dp.selectedDirection == None {
		return
	}

	if dp.selectedDirection == UpperLeft {
		dp.buttons[Upper].SelectButton(true)
		dp.buttons[Left].SelectButton(true)
	} else if dp.selectedDirection == UpperRight {
		dp.buttons[Upper].SelectButton(true)
		dp.buttons[Right].SelectButton(true)
	} else if dp.selectedDirection == LowerLeft {
		dp.buttons[Lower].SelectButton(true)
		dp.buttons[Left].SelectButton(true)
	} else if dp.selectedDirection == LowerRight {
		dp.buttons[Lower].SelectButton(true)
		dp.buttons[Right].SelectButton(true)
	} else {
		dp.buttons[dp.selectedDirection].SelectButton(true)
	}
}

// Draw draws the directional buttons belong this struct.
func (dp *DirectionalPad) Draw(screen *ebiten.Image) {
	screen.DrawImage(dp.baseImg, dp.op)
	for index := range dp.buttons {
		dp.buttons[index].Draw(screen)
	}
}

// GetDirection returns the currently selected direction.
func (dp *DirectionalPad) GetDirection() Direction {
	return dp.selectedDirection
}
