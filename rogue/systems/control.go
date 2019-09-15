package systems

import (
	"time"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/objects"
)

type InputProvider interface {
	GetDirection() cartography.Direction
}

type ControllerSystem struct {
	provider InputProvider
}

func NewControllerSystem(provider InputProvider) *ControllerSystem {
	return &ControllerSystem{
		provider: provider,
	}
}

func (c *ControllerSystem) ShouldTrack(object objects.GameObject) bool {
	return object.HasComponent(components.PlayerControlName) && object.HasComponent(components.PositionName)
}

func (c *ControllerSystem) Update(dT time.Duration, worldMap cartography.Map, objects map[uint64]objects.GameObject) error {
	dir := c.provider.GetDirection()
	for _, obj := range objects {
		obj.AddComponents(&components.Movement{Direction: dir})
	}
	return nil
}
