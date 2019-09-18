package components

import (
	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/purposed/good/datastructure/stringset"
)

const (
	CameraName = "camera"
)

// Camera represents an observer.
type Camera struct {
	SightRadius int
	Memory      cartography.Map

	BlockedBy stringset.StringSet
}

func (c *Camera) Name() string {
	return CameraName
}
