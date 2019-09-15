package components

import "github.com/dalloriam/rogue/rogue/cartography"

const (
	MovementName = "movement"
)

type Movement struct {
	Direction cartography.Direction
}

func (m *Movement) Name() string {
	return MovementName
}
