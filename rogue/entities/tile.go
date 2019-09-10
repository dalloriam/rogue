package entities

import (
	"image/color"

	"github.com/dalloriam/rogue/rogue/components"
)

// Tile represents a drawn tile.
type Tile struct {
	*BaseObject
}

// NewTile returns a new tile instance
func NewTile(posX, posY uint64, char rune, fgColor color.Color, bgColor color.Color) *Tile {
	return &Tile{
		BaseObject: NewObject(
			components.Drawable{
				Char:    char,
				FgColor: fgColor,
				BgColor: bgColor,
			},
			components.Position{
				X: posX,
				Y: posY,
			},
		),
	}
}
