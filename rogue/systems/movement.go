package systems

import (
	"time"

	"github.com/dalloriam/rogue/rogue/structure"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

type MovementSystem struct {
}

func NewMovementSystem() *MovementSystem {
	return &MovementSystem{}
}

func (s *MovementSystem) Name() string {
	return "movement"
}

func (s *MovementSystem) ShouldTrack(object object.GameObject) bool {
	return object.HasComponent(components.MovementName) && object.HasComponent(components.PositionName)
}

func (s *MovementSystem) Update(dT time.Duration, worldMap cartography.Map, objects map[uint64]object.GameObject) error {
	for _, obj := range objects {
		movement := obj.GetComponent(components.MovementName).(*components.Movement)
		position := obj.GetComponent(components.PositionName).(*components.Position)

		newPosition := components.Position{Vec: structure.V(position.X(), position.Y())}
		displacement := structure.V(0, 0)

		switch movement.Direction {
		case cartography.DirectionUp:
			displacement = structure.V(0, 1)
		case cartography.DirectionDown:
			displacement = structure.V(0, -1)
		case cartography.DirectionLeft:
			displacement = structure.V(-1, 0)
		case cartography.DirectionRight:
			displacement = structure.V(1, 0)
		case cartography.DirectionDownRight:
			displacement = structure.V(1, -1)
		case cartography.DirectionDownLeft:
			displacement = structure.V(-1, -1)
		case cartography.DirectionUpRight:
			displacement = structure.V(1, 1)
		case cartography.DirectionUpLeft:
			displacement = structure.V(-1, 1)
		}
		newPosition.Add(displacement)

		obj.RemoveComponent(components.MovementName)

		// Check if the movement is blocked before triggering.
		if obj.HasComponent(components.PhysicsName) {
			phys := obj.GetComponent(components.PhysicsName).(*components.Physics)
			tgtTile := worldMap.At(newPosition)

			if !phys.BlockedBy.Contains(tgtTile.Type) {
				obj.AddComponents(&newPosition)
			}
		} else {
			obj.AddComponents(&newPosition)
		}
	}
	return nil
}
