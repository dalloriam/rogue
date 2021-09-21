package components

import (
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
