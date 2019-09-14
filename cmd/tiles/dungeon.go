package tiles

import (
	"image/color"

	"github.com/dalloriam/rogue/rogue/cartography"
)

func RockWall(x, y int) cartography.Tile {
	return cartography.Tile{
		X:       x,
		Y:       y,
		Char:    '#',
		FgColor: color.Black,
		BgColor: color.Gray{128},
	}
}

func RockFloor(x, y int) cartography.Tile {
	return cartography.Tile{
		X:       x,
		Y:       y,
		Char:    '.',
		FgColor: color.White,
		BgColor: color.Black,
	}
}
