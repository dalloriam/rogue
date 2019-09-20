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
func (e *AutoExplorer) GetAction() func(obj object.GameObject) {
	d := []cartography.Direction{cartography.DirectionUp, cartography.DirectionLeft, cartography.DirectionRight, cartography.DirectionDown}

	return func(obj object.GameObject) {
		obj.AddComponents(&components.Movement{Direction: d[rand.Intn(len(d))]})
	}
}
