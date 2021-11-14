package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	vpad "local.packages/vpad"
)

const (
	screenWidth  = 480
	screenHeight = 320
)

func NewGame() *Game {
	dpad := vpad.NewDirectionalPad(directional_pad, directional_button, vpad.SelectColor)
	dpad.SetLocation(10, 180)

	// vpad.Pressing is continuously triggered while this button is being pressed.
	aButton := vpad.NewTriggerButton(a_button, vpad.Pressing, vpad.SelectColor)
	aButton.SetLocation(150, 190)

	// vpad.JustRelease is triggered when this button is released.
	bButton := vpad.NewTriggerButton(b_button, vpad.JustReleased, vpad.SelectColor)
	bButton.SetLocation(260, 190)

	// vpad.JustPress is triggered when this button is pressed.
	cButton := vpad.NewTriggerButton(c_button, vpad.JustPressed, vpad.SelectColor)
	cButton.SetLocation(370, 190)

	slButton := vpad.NewSelectButton(c_button, vpad.JustReleased, vpad.SelectColor)
	slButton.SetLocation(205, 90)

	srButton := vpad.NewSelectButton(c_button, vpad.JustPressed, vpad.SelectColor)
	srButton.SetLocation(315, 90)

	return &Game{
		dpad:                     dpad,
		pressingButton:           aButton,
		justReleasedButton:       bButton,
		justPressedButton:        cButton,
		justReleasedSelectButton: slButton,
		justPressedSelectButton:  srButton,
	}
}

type Game struct {
	dpad                     vpad.DirectionalPad
	pressingButton           vpad.TriggerButton
	justReleasedButton       vpad.TriggerButton
	justPressedButton        vpad.TriggerButton
	justReleasedSelectButton vpad.SelectButton
	justPressedSelectButton  vpad.SelectButton
}

func (g *Game) Update() error {
	g.dpad.Update()
	g.pressingButton.Update()
	g.justReleasedButton.Update()
	g.justPressedButton.Update()
	g.justReleasedSelectButton.Update()
	g.justPressedSelectButton.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.dpad.Draw(screen)
	d := g.dpad.GetDirection()
	switch d {
	case vpad.Upper:
		ebitenutil.DebugPrintAt(screen, "Upper direction", 20, 20)
	case vpad.Lower:
		ebitenutil.DebugPrintAt(screen, "Lower direction", 20, 20)
	case vpad.Left:
		ebitenutil.DebugPrintAt(screen, "Left direction", 20, 20)
	case vpad.Right:
		ebitenutil.DebugPrintAt(screen, "Right direction", 20, 20)
	}

	g.pressingButton.Draw(screen)
	if g.pressingButton.IsTriggered() {
		ebitenutil.DebugPrintAt(screen, "A button is being triggered.", 20, 50)
	}

	g.justReleasedButton.Draw(screen)
	if g.justReleasedButton.IsTriggered() {
		ebitenutil.DebugPrintAt(screen, "B button is triggered.", 20, 50)
	}

	g.justPressedButton.Draw(screen)
	if g.justPressedButton.IsTriggered() {
		ebitenutil.DebugPrintAt(screen, "C button is triggered.", 20, 50)
	}

	g.justReleasedSelectButton.Draw(screen)
	if g.justReleasedSelectButton.IsSelected() {
		ebitenutil.DebugPrintAt(screen, "L button is being selected.", 20, 70)
	}

	g.justPressedSelectButton.Draw(screen)
	if g.justPressedSelectButton.IsSelected() {
		ebitenutil.DebugPrintAt(screen, "R button is being selected.", 200, 70)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
