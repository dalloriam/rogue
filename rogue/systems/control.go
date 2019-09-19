package systems

import (
	"time"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

type InputProvider interface {
	GetDirection() cartography.Direction

	RepeatModeTriggered() bool
}

type ControllerSystem struct {
	provider InputProvider

	isInRepeatMode bool
}

func NewControllerSystem(provider InputProvider) *ControllerSystem {
	return &ControllerSystem{
		provider: provider,
	}
}

func (c *ControllerSystem) ShouldTrack(object object.GameObject) bool {
	return object.HasComponent(components.PlayerControlName) && object.HasComponent(components.PositionName)
}

func (c *ControllerSystem) Update(dT time.Duration, worldMap cartography.Map, objects map[uint64]object.GameObject) error {
	currentDirection := c.provider.GetDirection()
	for _, obj := range objects {
		// If a direction key is pressed, prioritize it and cancel repeat mode.
		if currentDirection != cartography.NoDirection {
			obj.AddComponents(&components.Movement{Direction: currentDirection})
			c.isInRepeatMode = false
		} else if c.provider.RepeatModeTriggered() {
			c.isInRepeatMode = true
		} else if c.isInRepeatMode {

		}
	}
	return nil
}
