package entities

import (
	"image/color"

	"github.com/dalloriam/rogue/backends/roguepixel"
	"github.com/dalloriam/rogue/rogue/ai"

	"github.com/dalloriam/rogue/rogue/structure"

	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
	"github.com/purposed/good/datastructure/stringset"
)

// Player builds a player object.
func Player(x, y int, handler *roguepixel.InputHandler) object.GameObject {
	return object.New(
		&components.Drawable{
			Char:    '@',
			FgColor: color.White,
			BgColor: color.RGBA{R: 0, G: 0, B: 0, A: 0},
		},
		&components.Position{
			Vec: structure.V(x, y),
		},
		&components.Physics{
			BlockedBy: stringset.FromValues([]string{"wall"}),
		},
		&components.Control{Agent: ai.NewPlayerController(handler)},
		&components.Focus{
			Priority: 0,
			Punctual: false,
		},
		&components.Camera{
			SightRadius: 8,
			BlockedBy:   stringset.FromValues([]string{"wall"}),
		},
		&components.Player{},
	)
}
