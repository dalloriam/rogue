package rogue

import (
	"sort"
	"time"

	"github.com/dalloriam/rogue/rogue/structure"

	"github.com/dalloriam/rogue/rogue/cartography"

	"github.com/dalloriam/rogue/rogue/systems"

	"github.com/dalloriam/rogue/rogue/object"
)

// World represents the root World.
type World struct {
	systemPriorities map[*systems.GameSystem]int
	systems          []*systems.GameSystem

	objects map[uint64]object.GameObject

	lastTick  time.Time
	turnClock *structure.TurnClock

	// worldMap represents the currently loaded map in its entirety -- NOT the cartography sections displayed in the viewport.
	worldMap cartography.Map
}

// NewWorld creates and returns a world.
func NewWorld() *World {
	return &World{
		systemPriorities: make(map[*systems.GameSystem]int),
		objects:          make(map[uint64]object.GameObject),
		lastTick:         time.Now(),
	}
}

// LoadMap loads a map into the world.
func (w *World) LoadMap(m cartography.Map) {
	w.worldMap = m
}

// AddObject adds an object to the list of objects managed by the world's systems.
func (w *World) AddObject(object object.GameObject) {
	// Add the object to the main registry.
	w.objects[object.ID()] = object
}

// AddSystem adds a new processing system to the world. The lower the priority, the later the system will run.
func (w *World) AddSystem(sys systems.System, priority int) {
	w.systemPriorities[systems.NewGameSystem(sys)] = priority

	// Re-sort the systems slice.
	// TODO: Optimize.
	var sysColl []*systems.GameSystem
	for system := range w.systemPriorities {
		sysColl = append(sysColl, system)
	}
	sort.Slice(sysColl, func(i, j int) bool {
		return w.systemPriorities[sysColl[i]] >= w.systemPriorities[sysColl[j]]
	})
	w.systems = sysColl
}

// Tick performs a single world update. Corresponds to a frame with non-cli renderers .
func (w *World) Tick() error {
	dT := time.Since(w.lastTick)
	w.lastTick = time.Now()

	for _, system := range w.systems {
		if err := system.Update(dT, w.worldMap, w.objects); err != nil {
			return err
		}
	}

	return nil
}
