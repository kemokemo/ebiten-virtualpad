package vpad

// TriggerType is the type for the TriggerButton.
type TriggerType int

const (
	// JustReleased is triggered only when just released.
	JustReleased TriggerType = iota
	// Pressing is triggered while this is being pressed.
	Pressing
	// JustPressed is triggered only when just pressed.
	JustPressed
)
