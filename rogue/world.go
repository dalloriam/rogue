package rogue

import (
	"sort"

	"github.com/dalloriam/rogue/rogue/systems"

	"github.com/dalloriam/rogue/rogue/entities"
)

// World represents the root World.
type World struct {
	systemPriorities map[*systems.GameSystem]int
	systems          []*systems.GameSystem

	objects map[uint64]entities.GameObject

	// currentMap represents the currently loaded map in its entirety -- NOT the map sections displayed in the viewport.
	currentMap Map
}

func NewWorld() *World {
	return &World{
		systemPriorities: make(map[*systems.GameSystem]int),
		objects:          make(map[uint64]entities.GameObject),
	}
}

func (w *World) LoadMap(m Map) {
	w.currentMap = m
}

func (w *World) AddObject(object entities.GameObject) {
	// Add the object to the main registry.
	w.objects[object.ID()] = object

	for _, system := range w.systems {
		system.AddObject(object)
	}
}

func (w *World) AddSystem(sys systems.System, priority int) {
	w.systemPriorities[systems.NewGameSystem(sys)] = priority

	// Re-sort the systems slice.
	// TODO: Optimize.
	var sysColl []*systems.GameSystem
	for system := range w.systemPriorities {
		sysColl = append(sysColl, system)
	}
	sort.Slice(sysColl, func(i, j int) bool {
		return w.systemPriorities[sysColl[i]] < w.systemPriorities[sysColl[j]]
	})
	w.systems = sysColl
}

func (w *World) Tick() error {
	for _, system := range w.systems {
		if err := system.Update(); err != nil {
			return err
		}
	}

	return nil
}
