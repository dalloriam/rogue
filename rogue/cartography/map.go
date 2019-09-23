package cartography

import (
	"fmt"
	"image/color"

	"github.com/dalloriam/rogue/rogue/structure"
)

type Tile struct {
	Position structure.Vec

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

func (m Map) At(position structure.Vec) *Tile {
	return &m[position.X()][position.Y()]
}

func (m Map) Set(tile Tile) {
	m[tile.Position.X()][tile.Position.Y()] = tile
}

func (m Map) aStarSearch(begin, end structure.Vec) (structure.Vec, float64) {
	frontier := make(PriorityQueue, 0)
	frontier.Push(&Item{Pos: begin, Score: 0})

	x := make(map[structure.Vec]structure.Vec)
	fmt.Println(x)

	return nil, 0
}

func (m Map) FindPath(begin, end structure.Vec) []structure.Vec {
	var path []structure.Vec

	return path
}
