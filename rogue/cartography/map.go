package cartography

import "image/color"

type Tile struct {
	X int
	Y int

	Type string // TODO: Find better way of representing tile type.

	Visibility float64 // 0.0 is invisible, 1.0 is fully lit.

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

func (m Map) At(x, y int) *Tile {
	return &m[x][y]
}

func (m Map) Set(tile Tile) {
	m[tile.X][tile.Y] = tile
}
