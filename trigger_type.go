package vpad

// TriggerType is the type for the TriggerButton.
type TriggerType int

const (
	// JustRelease is triggered only when just released.
	JustRelease TriggerType = iota
	// Pressing is triggered while this is being pressed.
	Pressing
	// JustPress is triggered only when just pressed.
	JustPress
)
