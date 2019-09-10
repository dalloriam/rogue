package systems

import "github.com/dalloriam/rogue/rogue/entities"

// System represents a system.
type System interface {
	// Abstract.
	ShouldTrack(object entities.GameObject) bool
	Update(objects map[uint64]entities.GameObject) error
}

// GameSystem wraps an abstract system in a well-composed system object.
type GameSystem struct {
	objects map[uint64]entities.GameObject
	system  System
}

// NewGameSystem returns a game system from the provided struct implementing the system interface.
func NewGameSystem(sys System) *GameSystem {
	return &GameSystem{
		objects: make(map[uint64]entities.GameObject),
		system:  sys,
	}
}

func (b *GameSystem) Update() error {
	return b.system.Update(b.objects)
}

func (b *GameSystem) initialize() {
	b.objects = make(map[uint64]entities.GameObject)
}

func (b *GameSystem) AddObject(object entities.GameObject) {
	if b.system.ShouldTrack(object) {
		b.objects[object.ID()] = object
	}
}

func (b *GameSystem) RemoveObject(id uint64) {
	delete(b.objects, id)
}