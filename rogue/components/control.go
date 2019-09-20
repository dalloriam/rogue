package components

import "github.com/dalloriam/rogue/rogue/object"

const (
	ControlName = "control"
)

type Agent interface {
	GetAction() func(obj object.GameObject)
}

type Control struct {
	Agent Agent
}

func (p *Control) Name() string {
	return ControlName
}
