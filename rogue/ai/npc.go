package ai

import (
	"fmt"
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

	if !obj.HasComponent(components.CameraName) {
		// This NPC is probably blind
		return func() {
			obj.AddComponents(&components.Movement{Direction: d[rand.Intn(len(d))]})
		}

	}

	return func() {
		cam := obj.GetComponent(components.CameraName).(*components.Camera)
		pos := obj.GetComponent(components.PositionName).(*components.Position)
		if len(cam.View) > 0 {
			// The camera view isn't initialized yet in the first frame - skip the sight check.
			// TODO: This is hacky as hell - improve.
			for _, o := range cam.View.ObjectsInRadius(pos, cam.SightRadius) {
				if o.ID() == obj.ID() {
					continue
				}
				fmt.Println("AI SAW OBJECT ", o.ID())
			}
		}
	}
}
