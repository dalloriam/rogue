package entities

import (
	"image/color"

	"github.com/dalloriam/rogue/rogue/components"
)

// Tile represents a drawn tile.
type Tile struct {
	*GameObject
}

// NewTile returns a new tile instance
func NewTile(char rune, fgColor color.Color, bgColor color.Color) *Tile {
	return &Tile{
		GameObject: NewObject(
			components.Drawable{
				Char:    char,
				FgColor: fgColor,
				BgColor: bgColor,
			},
			components.Position{
				X: 0,
				Y: 0,
			},
		),
	}
}
