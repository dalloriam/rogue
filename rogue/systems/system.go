package systems

import (
	"time"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/entities"
)

// System represents a system.
type System interface {
	// Abstract.
	ShouldTrack(object entities.GameObject) bool
	Update(dT time.Duration, worldMap cartography.Map, objects map[uint64]entities.GameObject) error
}

// GameSystem wraps an abstract system in a well-composed system object.
type GameSystem struct {
	system System
}

// NewGameSystem returns a game system from the provided struct implementing the system interface.
func NewGameSystem(sys System) *GameSystem {
	return &GameSystem{
		system: sys,
	}
}

func (b *GameSystem) Update(dT time.Duration, currentMap cartography.Map, objects map[uint64]entities.GameObject) error {
	// Filter out invalid entities.
	desiredObjects := make(map[uint64]entities.GameObject)
	for objID, obj := range objects {
		if b.system.ShouldTrack(obj) {
			desiredObjects[objID] = obj
		}
	}

	return b.system.Update(dT, currentMap, desiredObjects)
}
