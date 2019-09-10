package tiles

import (
	"image/color"

	"github.com/dalloriam/rogue/rogue/entities"
)

func RockWall(x, y uint64) *entities.Tile {
	return entities.NewTile(x, y, '#', color.Black, color.Gray{128})
}

func RockFloor(x, y uint64) *entities.Tile {
	return entities.NewTile(x, y, '.', color.White, color.Black)
}
