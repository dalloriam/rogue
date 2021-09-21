package ai

import (
	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

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
	repeatedAction func(obj object.GameObject)

	provider InputProvider
	state    inputProviderState
}

// NewPlayerController returns a new player controller.
func NewPlayerController(provider InputProvider) *PlayerController {
	return &PlayerController{
		provider: provider,
	}
}

func (pc *PlayerController) getPlayerInputAction(tgtObj object.GameObject) func(obj object.GameObject) {
	currentDirection := pc.provider.GetDirection()
	if currentDirection != cartography.NoDirection {
		return func(obj object.GameObject) {
			obj.AddComponents(&components.Movement{Direction: currentDirection})
		}
	}

	if pc.provider.RepeatModeTriggered() {
		pc.state = repeatWaitingForAction
		return nil
	}

	if pc.provider.AutoExploreTriggered() {
		tgtObj.AddComponents(&components.Control{Agent: NewAutoExplorer(pc)})
		return nil
	}
	return nil
}

// GetAction returns the action performed by this entity.
func (pc *PlayerController) GetAction(obj object.GameObject) func() {
	playerInputAction := pc.getPlayerInputAction(obj)

	if !obj.HasComponent(components.InitiativeName) {
		return nil
	}

	switch pc.state {
	case defaultState:
		if playerInputAction != nil {
			return func() { playerInputAction(obj) }
		}

	case repeatWaitingForAction:
		if playerInputAction != nil {
			pc.repeatedAction = playerInputAction
			pc.state = repeatMode
			return func() { playerInputAction(obj) }
		}
	case repeatMode:
		if playerInputAction == nil {
			return func() {
				if act := pc.getPlayerInputAction(obj); act != nil {
					obj.AddComponents(&components.Control{Agent: pc})
				}
				pc.repeatedAction(obj)
			}
		} else {
			pc.state = defaultState
		}
	}
	return nil
}
