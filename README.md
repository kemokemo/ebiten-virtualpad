# Ebiten Virtual Pad

This is a package for displaying virtual directional keys and trigger buttons on the screen to handle user input.  
It is intended to be used in games that use [ebiten](https://ebiten.org/).

## How to use

Please check the sample app [vpad-sample](https://github.com/kemokemo/vpad-sample).

### Basic

- Create: [here](https://github.com/kemokemo/vpad-sample/blob/0f957b01dc10ca8129e4202a8b4da5a4cb3fcb5a/game.go#L14)
- Update: [here](https://github.com/kemokemo/vpad-sample/blob/0f957b01dc10ca8129e4202a8b4da5a4cb3fcb5a/game.go#L36)
- Draw: [here](https://github.com/kemokemo/vpad-sample/blob/0f957b01dc10ca8129e4202a8b4da5a4cb3fcb5a/game.go#L44)

### Get direction and triggered state

- Direction: [here](https://github.com/kemokemo/vpad-sample/blob/0f957b01dc10ca8129e4202a8b4da5a4cb3fcb5a/game.go#L45)
- Triggered: [here](https://github.com/kemokemo/vpad-sample/blob/0f957b01dc10ca8129e4202a8b4da5a4cb3fcb5a/game.go#L58)

### Various triggers

- Triggered continuously while being pressed: [here](https://github.com/kemokemo/vpad-sample/blob/0f957b01dc10ca8129e4202a8b4da5a4cb3fcb5a/game.go#L18)
- Triggered once when released: [here](https://github.com/kemokemo/vpad-sample/blob/0f957b01dc10ca8129e4202a8b4da5a4cb3fcb5a/game.go#L22)

## Licence

Apache-2.0 License

## Author

kemokemo
