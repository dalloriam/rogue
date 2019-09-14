package cartography

import "image/color"

type Tile struct {
	X int
	Y int

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

func (m Map) At(x, y int) Tile {
	return m[x][y]
}

func (m Map) Set(x, y int, tile Tile) {
	m[x][y] = tile
}
