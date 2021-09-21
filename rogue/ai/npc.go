package ai

import (
	"math/rand"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

type NPCController struct{}

func NewNPCController() *NPCController {
	return &NPCController{}
}

func (npc *NPCController) GetAction(obj object.GameObject) func() {
	d := []cartography.Direction{
		cartography.DirectionUp,
		cartography.DirectionLeft,
		cartography.DirectionRight,
		cartography.DirectionDown,
	}

	return func() {
		obj.AddComponents(&components.Movement{Direction: d[rand.Intn(len(d))]})
	}

}
