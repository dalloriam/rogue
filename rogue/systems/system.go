package systems

import (
	"time"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

type UpdateInfo struct {
	DeltaT            time.Duration
	ObjectsByID       map[uint64]object.GameObject
	ObjectPositionMap [][][]uint64
	WorldMap          cartography.Map
}

// System represents a system.
type System interface {
	Name() string
	ShouldTrack(object object.GameObject) bool
	Update(info UpdateInfo) error
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

func (b *GameSystem) Name() string {
	return b.system.Name()
}

func (b *GameSystem) Update(info UpdateInfo) error {
	// Filter out invalid objects.
	desiredObjects := make(map[uint64]object.GameObject)
	for objID, obj := range info.ObjectsByID {
		if b.system.ShouldTrack(obj) {
			desiredObjects[objID] = obj
		}
	}

	objectMap := make([][][]uint64, len(info.ObjectPositionMap))
	for i := 0; i < len(info.ObjectPositionMap); i++ {
		objectMap[i] = make([][]uint64, len(info.ObjectPositionMap[i]))
	}
	for _, obj := range desiredObjects {
		if !obj.HasComponent(components.PositionName) {
			continue
		}

		position := obj.GetComponent(components.PositionName).(*components.Position)
		objectMap[position.X()][position.Y()] = append(objectMap[position.X()][position.Y()], obj.ID())
	}

	info.ObjectsByID = desiredObjects
	info.ObjectPositionMap = objectMap

	return b.system.Update(info)
}
