package ai

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

// InputProvider provides input from the player.
type InputProvider interface {
	GetDirection() cartography.Direction

	RepeatModeTriggered() bool
	AutoExploreTriggered() bool
}

// PlayerController exposes player behavior as an AI agent.
type PlayerController struct {
	lastActionRepeat time.Time
	repeatedAction   func(obj object.GameObject)

	provider InputProvider
	state    inputProviderState
}

// NewPlayerController returns a new player controller.
func NewPlayerController(provider InputProvider) *PlayerController {
	return &PlayerController{
		provider: provider,
	}
}

func (pc *PlayerController) getPlayerInputAction() func(obj object.GameObject) {
	currentDirection := pc.provider.GetDirection()
	if currentDirection != cartography.NoDirection {
		return func(obj object.GameObject) {
			obj.AddComponents(&components.Movement{Direction: currentDirection})
		}
	}

	if pc.provider.RepeatModeTriggered() {
		return func(obj object.GameObject) {
			pc.state = repeatWaitingForAction
		}
	}

	if pc.provider.AutoExploreTriggered() {
		return func(obj object.GameObject) {
			obj.AddComponents(&components.Control{Agent: NewAutoExplorer(pc)})
		}
	}
	return nil
}

// GetAction returns the action performed by this entity.
func (pc *PlayerController) GetAction(worldMap cartography.Map) func(obj object.GameObject) {
	playerInputAction := pc.getPlayerInputAction()

	switch pc.state {
	case defaultState:
		return playerInputAction

	case repeatWaitingForAction:
		if playerInputAction != nil {
			pc.repeatedAction = playerInputAction
			pc.state = repeatMode
			return playerInputAction
		}
	case repeatMode:
		if playerInputAction == nil && time.Since(pc.lastActionRepeat) > repeatActionMinDelta {
			pc.lastActionRepeat = time.Now()
			return pc.repeatedAction
		} else if playerInputAction != nil {
			pc.state = defaultState
			return playerInputAction
		}
	}
	return nil
}
