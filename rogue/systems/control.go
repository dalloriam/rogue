package systems

import (
	"time"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

type ControllerSystem struct{}

func NewControllerSystem() *ControllerSystem {
	return &ControllerSystem{}
}

// ShouldTrack returns whether
func (c *ControllerSystem) ShouldTrack(object object.GameObject) bool {
	return object.HasComponent(components.ControlName) && object.HasComponent(components.PositionName)
}

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
