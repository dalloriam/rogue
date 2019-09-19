package systems

import (
	"time"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

type MovementSystem struct {
}

func NewMovementSystem() *MovementSystem {
	return &MovementSystem{}
}

func (s *MovementSystem) ShouldTrack(object object.GameObject) bool {
	return object.HasComponent(components.MovementName) && object.HasComponent(components.PositionName)
}

func (s *MovementSystem) Update(dT time.Duration, worldMap cartography.Map, objects map[uint64]object.GameObject) error {
	for _, object := range objects {
		movement := object.GetComponent(components.MovementName).(*components.Movement)
		position := object.GetComponent(components.PositionName).(*components.Position)

		newPosition := components.Position{
			X: position.X,
			Y: position.Y,
		}

		switch movement.Direction {
		case cartography.DirectionUp:
			newPosition.Y++
		case cartography.DirectionDown:
			newPosition.Y--
		case cartography.DirectionLeft:
			newPosition.X--
		case cartography.DirectionRight:
			newPosition.X++
		}

		object.RemoveComponent(components.MovementName)

		// Check if the movement is blocked before triggering.
		if object.HasComponent(components.PhysicsName) {
			phys := object.GetComponent(components.PhysicsName).(*components.Physics)
			tgtTile := worldMap.At(newPosition.X, newPosition.Y)

			if !phys.BlockedBy.Contains(tgtTile.Type) {
				object.AddComponents(&newPosition)
			}
		} else {
			object.AddComponents(&newPosition)
		}
	}
	return nil
}
