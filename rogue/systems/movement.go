package systems

import (
	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/entities"
)

type MovementSystem struct {
}

func NewMovementSystem() *MovementSystem {
	return &MovementSystem{}
}

func (s *MovementSystem) ShouldTrack(object entities.GameObject) bool {
	return object.HasComponent(components.MovementName) && object.HasComponent(components.PositionName)
}

func (s *MovementSystem) Update(worldMap cartography.Map, objects map[uint64]entities.GameObject) error {
	for _, object := range objects {
		movement := object.GetComponent(components.MovementName).(components.Movement)
		position := object.GetComponent(components.PositionName).(components.Position)

		switch movement.Direction {
		case components.DirectionUp:
			position.Y++
		case components.DirectionDown:
			position.Y--
		case components.DirectionLeft:
			position.X--
		case components.DirectionRight:
			position.X++
		}

		object.RemoveComponent(components.MovementName)
	}
	return nil
}
