package components

import (
	"github.com/dalloriam/rogue/rogue/structure"
	"github.com/purposed/good/datastructure/stringset"
)

// Name of the component
const (
	CameraName = "camera"
)

// Camera represents an observer.
type Camera struct {
	SightRadius int
	Memory      []structure.Vec

	BlockedBy stringset.StringSet
}

// Name returns the name of the component.
func (c *Camera) Name() string {
	return CameraName
}
