package components

import "github.com/purposed/good/datastructure/stringset"

const (
	PhysicsName = "physics"
)

type Physics struct {
	BlockedBy stringset.StringSet
}

func (p *Physics) Name() string {
	return PhysicsName
}
