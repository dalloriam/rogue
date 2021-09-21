package rogue

import (
	"sort"
	"time"

	"github.com/dalloriam/rogue/rogue/components"
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

func NewWorld() *World {
	return &World{
		systemPriorities: make(map[*systems.GameSystem]int),
		objects:          make(map[uint64]object.GameObject),
		lastTick:         time.Now(),
	}
}

func (w *World) LoadMap(m cartography.Map) {
	w.worldMap = m
}

func (w *World) AddObject(object object.GameObject) {
	// Add the object to the main registry.
	w.objects[object.ID()] = object
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
		return w.systemPriorities[sysColl[i]] >= w.systemPriorities[sysColl[j]]
	})
	w.systems = sysColl
}

func (w *World) computeObjectPositionalMap() [][][]uint64 {
	objectMap := make([][][]uint64, w.worldMap.SizeX())
	for i := 0; i < w.worldMap.SizeX(); i++ {
		objectMap[i] = make([][]uint64, w.worldMap.SizeY())
	}
	for _, obj := range w.objects {
		if !obj.HasComponent(components.PositionName) {
			continue
		}

		position := obj.GetComponent(components.PositionName).(*components.Position)
		objectMap[position.X()][position.Y()] = append(objectMap[position.X()][position.Y()], obj.ID())
	}

	return objectMap
}

func (w *World) Tick() error {
	dT := time.Since(w.lastTick)
	w.lastTick = time.Now()

	info := systems.UpdateInfo{
		DeltaT:            dT,
		ObjectsByID:       w.objects,
		ObjectPositionMap: w.computeObjectPositionalMap(),
		WorldMap:          w.worldMap,
	}

	for _, system := range w.systems {
		if err := system.Update(info); err != nil {
			return err
		}
	}

	return nil
}
