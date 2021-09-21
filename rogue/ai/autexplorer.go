package ai

import (
	"math/rand"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

// AutoExplorer automatically explores the map, allowing the player to break out if he wants.
type AutoExplorer struct {
	pc *PlayerController
}

// NewAutoExplorer initializes & returns an explorer entity.
func NewAutoExplorer(pc *PlayerController) *AutoExplorer {
	return &AutoExplorer{pc: pc}
}

// GetAction returns this player's action.
func (e *AutoExplorer) GetAction(obj object.GameObject) func() {
	if act := e.pc.getPlayerInputAction(obj); act != nil {
		obj.AddComponents(&components.Control{Agent: e.pc})
		return nil
	}
	d := []cartography.Direction{cartography.DirectionUp, cartography.DirectionLeft, cartography.DirectionRight, cartography.DirectionDown}

	return func() {
		if act := e.pc.getPlayerInputAction(obj); act != nil {
			obj.AddComponents(&components.Control{Agent: e.pc})
		}
		obj.AddComponents(&components.Movement{Direction: d[rand.Intn(len(d))]})
	}
}
