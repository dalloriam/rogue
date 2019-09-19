package entities

import (
	"image/color"

	"github.com/dalloriam/rogue/rogue/structure"

	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
	"github.com/purposed/good/datastructure/stringset"
)

func Player(x, y int) object.GameObject {
	return object.New(
		&components.Drawable{
			Char:    '@',
			FgColor: color.White,
			BgColor: color.RGBA{0, 0, 0, 0},
		},
		&components.Position{
			Vec: structure.V(x, y),
		},
		&components.Physics{
			BlockedBy: stringset.FromValues([]string{"wall"}),
		},
		&components.PlayerControl{},
		&components.Focus{
			Priority: 0,
			Punctual: false,
		},
		&components.Camera{
			SightRadius: 8,
			BlockedBy:   stringset.FromValues([]string{"wall"}),
		},
	)
}
