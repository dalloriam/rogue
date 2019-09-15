package components

const (
	PhysicsName = "physics"
)

type Physics struct {
	BlockedBy []string
}

func (p *Physics) Name() string {
	return PhysicsName
}

func (p *Physics) IsBlocked(tileType string) bool {
	for _, t := range p.BlockedBy {
		if t == tileType {
			return true
		}
	}
	return false
}
