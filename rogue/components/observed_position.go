package components

import "github.com/dalloriam/rogue/rogue/structure"

const (
	ObservedPositionName = "observed_position"
)

type ObservedPosition struct {
	structure.Vec
}

func (p *ObservedPosition) Name() string {
	return ObservedPositionName
}
