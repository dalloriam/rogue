package systems

import (
	"time"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

var repeatActionMinDelta = 100 * time.Millisecond

type inputProviderState int

const (
	defaultState inputProviderState = iota
	repeatWaitingForAction
	repeatMode
)

type InputProvider interface {
	GetDirection() cartography.Direction

	RepeatModeTriggered() bool
}

type ControllerSystem struct {
	provider InputProvider

	state inputProviderState

	repeatedAction   action
	lastActionRepeat time.Time
}

type action func(obj object.GameObject)

func NewControllerSystem(provider InputProvider) *ControllerSystem {
	return &ControllerSystem{
		provider: provider,
	}
}

func (c *ControllerSystem) ShouldTrack(object object.GameObject) bool {
	return object.HasComponent(components.PlayerControlName) && object.HasComponent(components.PositionName)
}

func (c *ControllerSystem) getAction() action {
	currentDirection := c.provider.GetDirection()
	if currentDirection != cartography.NoDirection {
		return func(obj object.GameObject) {
			obj.AddComponents(&components.Movement{Direction: currentDirection})
		}
	}

	if c.provider.RepeatModeTriggered() {
		return func(obj object.GameObject) {
			if c.state == defaultState {
				c.state = repeatWaitingForAction
				c.lastActionRepeat = time.Now()
			} else {
				c.state = defaultState
			}
		}
	}

	return nil
}

func (c *ControllerSystem) Update(dT time.Duration, worldMap cartography.Map, objects map[uint64]object.GameObject) error {
	playerAction := c.getAction()

	for _, obj := range objects {
		switch c.state {
		case repeatMode:
			if playerAction == nil && time.Since(c.lastActionRepeat) > repeatActionMinDelta {
				c.repeatedAction(obj)
				c.lastActionRepeat = time.Now()
			} else if playerAction != nil {
				c.state = defaultState
				playerAction(obj)
			}
		case repeatWaitingForAction:
			if playerAction != nil {
				c.repeatedAction = playerAction
				c.state = repeatMode
			}
		case defaultState:
			if playerAction != nil {
				playerAction(obj)
			}
		}
	}
	return nil
}
