package entities

import (
	"image/color"

	"github.com/dalloriam/rogue/rogue/ai"

	"github.com/dalloriam/rogue/rogue/structure"

	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
	"github.com/purposed/good/datastructure/stringset"
)

func NPC(x, y int) object.GameObject {
	return object.New(
		&components.Drawable{
			Char:    'F',
			FgColor: color.RGBA{R: 255, G: 0, B: 0, A: 0},
			BgColor: color.RGBA{R: 0, G: 0, B: 0, A: 0},
		},
		&components.Position{
			Vec: structure.V(x, y),
		},
		&components.Physics{
			BlockedBy: stringset.FromValues([]string{"wall"}),
		},
		&components.Camera{
			SightRadius: 3,
			BlockedBy:   stringset.FromValues([]string{"wall"}),
		},
		&components.Control{Agent: ai.NewNPCController()},
		&components.Player{},
	)
}
