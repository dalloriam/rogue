package components

import "github.com/purposed/good/datastructure/stringset"

// Name of this component
const (
	PhysicsName = "physics"
)

// The Physics component indicates that the object is affected by physics.
type Physics struct {
	BlockedBy stringset.StringSet
}

// Name returns the component's name.
func (p *Physics) Name() string {
	return PhysicsName
}
