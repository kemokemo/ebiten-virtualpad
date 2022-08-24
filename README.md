# Ebiten Virtual Pad

This is a package for displaying virtual directional keys and trigger buttons on the screen to handle user input.  
It is intended to be used in games that use [ebiten](https://ebiten.org/).

## How to use

Please check [the sample app](https://github.com/kemokemo/ebiten-virtualpad/tree/main/cmd) for more detail.

### Basic

Use the `NewDirectionalPad` and `NewTriggerButton` functions of the vpad package to generate the buttons.

```go
import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	vpad "github.com/kemokemo/ebiten-virtualpad"
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

	return &Game{dpad: dpad, aButton: aButton, bButton: bButton, cButton: cButton}
}

type Game struct {
	dpad    vpad.DirectionalPad
	aButton vpad.TriggerButton
	bButton vpad.TriggerButton
	cButton vpad.TriggerButton
}
```

### Get direction and triggered state

Update the internal state with the `Update` function, and get the direction or trigger presence with the `GetDirection` or `IsTriggered` function.

```go
func (g *Game) Update() error {
	g.dpad.Update()
	g.aButton.Update()
	g.bButton.Update()
	g.cButton.Update()

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

	g.aButton.Draw(screen)
	if g.aButton.IsTriggered() {
		ebitenutil.DebugPrintAt(screen, "A button is being triggered.", 20, 50)
	}

	g.bButton.Draw(screen)
	if g.bButton.IsTriggered() {
		ebitenutil.DebugPrintAt(screen, "B button is triggered.", 20, 50)
	}

	g.cButton.Draw(screen)
	if g.cButton.IsTriggered() {
		ebitenutil.DebugPrintAt(screen, "C button is triggered.", 20, 50)
	}
}
```

### Various triggers

The following types of TriggerButtons are currently available.

| Type  | Description  |
|---|---|
| JustReleased  | This is triggered only when just released.  |
|Pressing|This is triggered while this is being pressed.|
|JustPressed|This is triggered only when just pressed.|

## License

Apache-2.0 License

## Author

kemokemo
