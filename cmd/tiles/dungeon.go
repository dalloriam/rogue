package tiles

import (
	"image/color"

	"github.com/dalloriam/rogue/rogue/structure"

	"github.com/dalloriam/rogue/rogue/cartography"
)

func RockWall(x, y int) cartography.Tile {
	return cartography.Tile{
		Position: structure.V(x, y),
		Char:     '#',
		FgColor:  color.Black,
		BgColor:  color.Gray{Y: 128},
		Type:     "wall",
	}
}

func RockFloor(x, y int) cartography.Tile {
	return cartography.Tile{
		Position: structure.V(x, y),
		Char:     '.',
		FgColor:  color.White,
		BgColor:  color.Black,
		Type:     "floor",
	}
}
