package systems

import (
	"image/color"
	"math"
	"time"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

// RenderingEngine abstracts a rendering engine.
type RenderingEngine interface {
	Clear()
	Draw()

	Rectangle(x, y int, bgColor color.Color)
	Text(x, y int, text string, fgColor color.Color)
}

// A Renderer renders components.
type Renderer struct {
	engine RenderingEngine
}

// NewRenderer returns a new rendering system.
func NewRenderer(engine RenderingEngine) (*Renderer, error) {
	return &Renderer{
		engine: engine,
	}, nil
}

// ShouldTrack returns true if the object has a position component & is drawable.
func (r *Renderer) ShouldTrack(object object.GameObject) bool {
	return object.HasComponent(components.DrawableName) && object.HasComponent(components.PositionName)
}

func (r *Renderer) clip(val, lowerBound, upperBound float64) float64 {
	if val > upperBound {
		return upperBound
	} else if val < lowerBound {
		return lowerBound
	}
	return val
}

func (r *Renderer) shadeColor(c color.Color, shadePercent float64) color.Color {
	red, g, b, a := c.RGBA()

	newR := r.clip(float64(red)*shadePercent, 0, math.MaxUint16)
	newG := r.clip(float64(g)*shadePercent, 0, math.MaxUint16)
	newB := r.clip(float64(b)*shadePercent, 0, math.MaxUint16)

	return color.RGBA64{
		R: uint16(newR),
		G: uint16(newG),
		B: uint16(newB),
		A: uint16(a),
	}
}

// Update updates the system state.
func (r *Renderer) Update(dT time.Duration, worldMap cartography.Map, objects map[uint64]object.GameObject) error {

	// TODO: Get rid of this.
	// Phase 0 - Clear previous frame.
	r.engine.Clear()

	// - Phase 1 -
	// Create an object positional map. This will help us drawing the map
	// by allowing us to perform drawable object lookups by position in O(1).
	objectMap := make([][]uint64, worldMap.SizeX())
	for i := 0; i < worldMap.SizeX(); i++ {
		objectMap[i] = make([]uint64, worldMap.SizeY())
	}
	for _, obj := range objects {
		position := obj.GetComponent(components.PositionName).(*components.Position)
		objectMap[position.X()][position.Y()] = obj.ID()
	}

	// Phase 1 - Draw cartography changes (TODO: Take viewport into account?)
	for i := 0; i < len(worldMap); i++ {
		for j := 0; j < len(worldMap[i]); j++ {
			currentTile := worldMap[i][j]
			if currentTile.Visibility == 0.0 {
				continue
			}

			// When drawing the tiles initially, we have no clue if we have an entity at this position.
			// First, what we know for sure is that we need to draw the map tile background.
			r.engine.Rectangle(i, j, r.shadeColor(currentTile.BgColor, currentTile.Visibility))

			// Once we have the tile background, we need to check if we have an object on this tile.
			// If so, we'll draw the object (Bg & Fg), otherwise we'll draw the tile foreground.
			if objectMap[currentTile.Position.X()][currentTile.Position.Y()] == 0 {
				// We don't have an entity. Proceed with the foreground
				r.engine.Text(i, j, string([]rune{currentTile.Char}), r.shadeColor(currentTile.FgColor, currentTile.Visibility))
			} else {
				// We have an object. First, draw its background (if it's not transparent).
				// TODO: Perform this check *before* rendering the tile background to save a drawing call.
				drawable := objects[objectMap[i][j]].GetComponent(components.DrawableName).(*components.Drawable)
				r.engine.Rectangle(i, j, r.shadeColor(drawable.BgColor, currentTile.Visibility))
				r.engine.Text(i, j, string([]rune{drawable.Char}), r.shadeColor(drawable.FgColor, currentTile.Visibility))
			}
		}
	}

	// Phase 2 - Commit to screen
	r.engine.Draw()

	return nil
}
