package cartography

import (
	"image/color"

	"github.com/dalloriam/rogue/rogue/structure"
)

type TileVisibility float64

const (
	VisibilityNone       TileVisibility = 0.0
	VisibilityOutOfSight TileVisibility = 0.25
	VisibilityInSight                   = 1.0
)

type Tile struct {
	Position structure.Vec

	Type string // TODO: Find better way of representing tile type.

	Visibility TileVisibility

	Char    rune
	FgColor color.Color
	BgColor color.Color
}

type Map [][]Tile

func NewMap(x, y int) Map {
	m := make([][]Tile, x)

	for i := 0; i < x; i++ {
		m[i] = make([]Tile, y)
	}

	return m
}

func (m Map) SizeX() int {
	return len(m)
}

func (m Map) SizeY() int {
	if len(m) == 0 {
		return 0
	}

	return len(m[0])
}

func (m Map) At(position structure.Vec) *Tile {
	return &m[position.X()][position.Y()]
}

func (m Map) Set(tile Tile) {
	m[tile.Position.X()][tile.Position.Y()] = tile
}

func (m Map) Copy() Map {
	n := NewMap(m.SizeX(), m.SizeY())

	for i := 0; i < m.SizeX(); i++ {
		for j := 0; j < m.SizeY(); j++ {
			n.Set(m[i][j])
			n[i][j] = m[i][j]
		}
	}

	return n
}
