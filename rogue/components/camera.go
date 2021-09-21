package components

import (
	"fmt"
	"math"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/object"
	"github.com/dalloriam/rogue/rogue/structure"
	"github.com/purposed/good/datastructure/stringset"
)

const (
	CameraName = "camera"
)

type ViewInfo struct {
	Tile     cartography.Tile
	Entities []object.GameObject
}

type CameraView [][]ViewInfo

func (c CameraView) At(position structure.Vec) *ViewInfo {
	return &c[position.X()][position.Y()]
}

func (c CameraView) ObjectsInRadius(position structure.Vec, sightRange int) []object.GameObject {
	// TODO: Fix this and make it a circular radius instead of a square one.
	objects := []object.GameObject{}

	xMin := int(math.Max(float64(position.X()-int(sightRange)), 0))
	xMax := int(math.Min(float64(position.X()+int(sightRange)), float64(len(c))))

	yMin := int(math.Max(float64(position.Y()-int(sightRange)), 0))
	yMax := int(math.Min(float64(position.Y()+int(sightRange)), float64(len(c[0]))))

	for i := xMin; i < xMax; i++ {
		for j := yMin; j < yMax; j++ {
			objects = append(objects, c[i][j].Entities...)
		}
	}

	fmt.Printf("%d objects in radius\n", len(objects))

	return objects
}

// Camera represents an observer.
type Camera struct {
	SightRadius int
	View        CameraView
	Main        bool

	BlockedBy stringset.StringSet
}

func (c *Camera) Name() string {
	return CameraName
}
