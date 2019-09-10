package components

const (
	PositionName = "position"
)

// Position represents a X, Y position in the window
type Position struct {
	X uint64
	Y uint64
}

func (p Position) Name() string {
	return PositionName
}
