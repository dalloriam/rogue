package tiles

import (
	"image/color"

	"github.com/dalloriam/rogue/rogue/structure"

	"github.com/dalloriam/rogue/rogue/cartography"
)

// RockWall creates a rock wall tile.
func RockWall(x, y int) cartography.Tile {
	return cartography.Tile{
		Position: structure.V(x, y),
		Char:     '#',
		FgColor:  color.Black,
		BgColor:  color.Gray{Y: 128},
		Type:     "wall",
	}
}

// RockFloor creates a rock floor walkable tile.
func RockFloor(x, y int) cartography.Tile {
	return cartography.Tile{
		Position: structure.V(x, y),
		Char:     '.',
		FgColor:  color.White,
		BgColor:  color.Black,
		Type:     "floor",
	}
}
