package components

import "github.com/dalloriam/rogue/rogue/structure"

// Name of the component.
const (
	PositionName = "position"
)

// The Position component indicates the object's location in the current map.
type Position struct {
	structure.Vec
}

// Name returns the component's name.
func (p *Position) Name() string {
	return PositionName
}
