package cartography

import (
	"image/color"

	"github.com/dalloriam/rogue/rogue/structure"
)

// A Tile is a grid element.
type Tile struct {
	Position structure.Vec

	Type string // TODO: Find better way of representing tile type.

	Visibility float64 // 0.0 is invisible, 1.0 is fully lit.

	Char    rune
	FgColor color.Color
	BgColor color.Color
}

// A Map is a 2-D array of tiles.
type Map [][]Tile

// NewMap returns a map with the specified dimensions.
func NewMap(x, y int) Map {
	m := make([][]Tile, x)

	for i := 0; i < x; i++ {
		m[i] = make([]Tile, y)
	}

	return m
}

// SizeX returns the horizontal map size.
func (m Map) SizeX() int {
	return len(m)
}

// SizeY returns the vertical map size.
func (m Map) SizeY() int {
	if len(m) == 0 {
		return 0
	}

	return len(m[0])
}

// At returns a reference to the tile at this specific position.
func (m Map) At(position structure.Vec) *Tile {
	return &m[position.X()][position.Y()]
}

// Set overwrites a specific tile.
func (m Map) Set(tile Tile) {
	m[tile.Position.X()][tile.Position.Y()] = tile
}
