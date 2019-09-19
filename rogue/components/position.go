package components

import "github.com/dalloriam/rogue/rogue/structure"

const (
	PositionName = "position"
)

// Position represents a X, Y position in the window
type Position struct {
	structure.Vec
}

func (p *Position) Name() string {
	return PositionName
}
