package components

const (
	PositionName = "position"
)

// Position represents a X, Y position in the window
type Position struct {
	X int64
	Y int64
}

func (p Position) Name() string {
	return PositionName
}
