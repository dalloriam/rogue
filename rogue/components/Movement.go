package components

const (
	MovementName = "movement"
)

type Direction string

const (
	DirectionUp    Direction = "up"
	DirectionDown  Direction = "down"
	DirectionLeft  Direction = "left"
	DirectionRight Direction = "right"
)

type Movement struct {
	Direction Direction
}

func (m Movement) Name() string {
	return MovementName
}
