package components

import (
	"github.com/dalloriam/rogue/rogue/object"
)

const (
	ControlName = "control"
)

type Agent interface {
	GetAction(obj object.GameObject) func()
}

type Control struct {
	Agent Agent
}

func (p *Control) Name() string {
	return ControlName
}
