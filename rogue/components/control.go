package components

import (
	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/object"
)

const (
	ControlName = "control"
)

type Agent interface {
	GetAction(obj object.GameObject, worldMap cartography.Map) func()
}

type Control struct {
	Agent Agent
}

func (p *Control) Name() string {
	return ControlName
}
