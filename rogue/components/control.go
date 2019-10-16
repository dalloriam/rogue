package components

import (
	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/object"
)

// Name of the component.
const (
	ControlName = "control"
)

// An Agent represents an entity making decisions for a controllable game object.
type Agent interface {
	GetAction(obj object.GameObject, worldMap cartography.Map) func()
}

// The Control component makes an object entity-controllable.
type Control struct {
	Agent Agent
}

// Name returns the name of the component.
func (p *Control) Name() string {
	return ControlName
}
