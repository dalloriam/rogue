package ai

import (
	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/object"
)

// AutoExplorer automatically explores the map, allowing the player to break out if he wants.
type AutoExplorer struct {
	pc PlayerController
}

// NewAutoExplorer initializes & returns an explorer entity.
func NewAutoExplorer(pc PlayerController) *AutoExplorer {
	return &AutoExplorer{pc: pc}
}

// GetAction returns this player's action.
func (e *AutoExplorer) GetAction(obj object.GameObject, worldMap cartography.Map) func(obj object.GameObject) {
	return nil
}
