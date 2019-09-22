package components

import (
	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/object"
)

const (
	ControlName = "control"
)

type Agent interface {
	GetAction(worldMap cartography.Map) func(obj object.GameObject)
}

type Control struct {
	Agent Agent
}

func (p *Control) Name() string {
	return ControlName
}
