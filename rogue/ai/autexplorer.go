package ai

import (
	"math/rand"

	"github.com/dalloriam/rogue/rogue/structure"

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

func (e *AutoExplorer) getUnexploredTiles(worldMap cartography.Map) []*cartography.Tile {
	var tiles []*cartography.Tile

	for i := 0; i < worldMap.SizeX(); i++ {
		for j := 0; j < worldMap.SizeY(); j++ {
			if t := worldMap.At(structure.V(i, j)); t.Visibility == 0.0 {
				// Tile is unseen.
				tiles = append(tiles, t)
			}
		}
	}

	return tiles
}

// GetAction returns this player's action.
func (e *AutoExplorer) GetAction(obj object.GameObject, worldMap cartography.Map) func() {
	if act := e.pc.getPlayerInputAction(obj); act != nil {
		obj.AddComponents(&components.Control{Agent: e.pc})
		return nil
	}

	unexplored := e.getUnexploredTiles(worldMap)
	// TODO: Use heuristic to sort tiles by most worthwhile of exploration.
	d := []cartography.Direction{cartography.DirectionUp, cartography.DirectionLeft, cartography.DirectionRight, cartography.DirectionDown}

	return func() {
		obj.AddComponents(&components.Movement{Direction: d[rand.Intn(len(d))]})
	}
}
