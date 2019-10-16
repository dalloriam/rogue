package components

import "github.com/dalloriam/rogue/rogue/cartography"

// Name of the component.
const (
	MovementName = "movement"
)

// The Movement component represents a movement applied to the object for this frame.
type Movement struct {
	Direction cartography.Direction
}

// Name returns the component's name.
func (m *Movement) Name() string {
	return MovementName
}
