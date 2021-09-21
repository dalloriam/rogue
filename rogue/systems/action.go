package systems

import (
	"time"

	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

var minTurnDelta = 100 * time.Millisecond // TODO: Make configurable.

type ActionSystem struct {
	timeOfLastTurn time.Time
}

func NewActionSystem() *ActionSystem {
	return &ActionSystem{
		timeOfLastTurn: time.Unix(0, 0),
	}
}

func (c *ActionSystem) Name() string {
	return "action"
}

// ShouldTrack returns whether
func (c *ActionSystem) ShouldTrack(object object.GameObject) bool {
	return object.HasComponent(components.ControlName) && object.HasComponent(components.PositionName)
}

func (c *ActionSystem) Update(info UpdateInfo) error {
	if time.Since(c.timeOfLastTurn) < minTurnDelta {
		// No action -- not enough time.
		return nil
	}

	actions := []func(){}

	for _, obj := range info.ObjectsByID {
		control := obj.GetComponent(components.ControlName).(*components.Control)
		action := control.Agent.GetAction(obj)

		if obj.HasComponent(components.InitiativeName) {
			if action == nil {
				// We're waiting for someone to take their turn and they're not moving now.
				return nil
			} else {
				actions = append(actions, action)
				obj.RemoveComponent(components.InitiativeName)
			}
		} else if action != nil {
			actions = append(actions, action)
		}
	}

	c.timeOfLastTurn = time.Now()

	for _, action := range actions {
		action()
	}
	return nil
}
