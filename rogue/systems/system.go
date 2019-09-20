package systems

import (
	"time"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/object"
)

// System represents a system.
type System interface {
	ShouldTrack(object object.GameObject) bool
	Update(dT time.Duration, worldMap cartography.Map, objects map[uint64]object.GameObject) error
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

func (b *GameSystem) Update(dT time.Duration, currentMap cartography.Map, gameObjects map[uint64]object.GameObject) error {
	// Filter out invalid objects.
	desiredObjects := make(map[uint64]object.GameObject)
	for objID, obj := range gameObjects {
		if b.system.ShouldTrack(obj) {
			desiredObjects[objID] = obj
		}
	}

	return b.system.Update(dT, currentMap, desiredObjects)
}
