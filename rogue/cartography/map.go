package cartography

import "image/color"

type Tile struct {
	X uint64
	Y uint64

	Char    rune
	FgColor color.Color
	BgColor color.Color
}

type Map [][]Tile

func NewMap(x, y uint64) Map {
	m := make([][]Tile, x)

	var i uint64
	for i = 0; i < x; i++ {
		m[i] = make([]Tile, y)
	}

	return m
}

func (m Map) SizeX() uint64 {
	return uint64(len(m))
}

func (m Map) SizeY() uint64 {
	if len(m) == 0 {
		return 0
	}

	return uint64(len(m[0]))
}

func (m Map) At(x, y uint64) Tile {
	return m[x][y]
}

func (m Map) Set(x, y uint64, tile Tile) {
	m[x][y] = tile
}
