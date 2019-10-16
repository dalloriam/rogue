package systems

import (
	"time"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

// The ControllerSystem applies actions performed by the various game objects.
type ControllerSystem struct{}

// NewControllerSystem returns an empty controller system.
func NewControllerSystem() *ControllerSystem {
	return &ControllerSystem{}
}

// ShouldTrack returns whether
func (c *ControllerSystem) ShouldTrack(object object.GameObject) bool {
	return object.HasComponent(components.ControlName) && object.HasComponent(components.PositionName)
}

// Update updates the system for a tick.
func (c *ControllerSystem) Update(dT time.Duration, worldMap cartography.Map, objects map[uint64]object.GameObject) error {
	for _, obj := range objects {
		control := obj.GetComponent(components.ControlName).(*components.Control)
		if action := control.Agent.GetAction(obj, worldMap); action != nil && obj.HasComponent(components.InitiativeName) {
			action()
			obj.RemoveComponent(components.InitiativeName)
		}
	}
	return nil
}
