package systems

import (
	"image/color"
	"time"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

// RenderingBackend abstracts a rendering engine.
type RenderingEngine interface {
	Clear()
	Draw()

	Rectangle(startX, startY, endX, endY int, bgColor color.Color)
	Text(startX, startY int, text string, fgColor color.Color)
}

type RendererOptions struct {
	TileSizeX int
	TileSizeY int
}

// A Renderer renders components.
type Renderer struct {
	engine RenderingEngine
	opt    RendererOptions
}

// NewRenderer returns a new rendering system.
func NewRenderer(engine RenderingEngine, opt RendererOptions) (*Renderer, error) {
	return &Renderer{
		engine: engine,
		opt:    opt,
	}, nil
}

func (r *Renderer) ShouldTrack(object object.GameObject) bool {
	return object.HasComponent(components.DrawableName) && object.HasComponent(components.PositionName)
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
		objectMap[position.X][position.Y] = obj.ID()
	}

	// Phase 1 - Draw cartography changes (TODO: Take viewport into account?)
	for i := 0; i < len(worldMap); i++ {
		for j := 0; j < len(worldMap[i]); j++ {
			currentTile := worldMap[i][j]
			// When drawing the tiles initially, we have no clue if we have an entity at this position.
			// First, what we know for sure is that we need to draw the map tile background.
			startX := currentTile.X * r.opt.TileSizeX
			startY := currentTile.Y * r.opt.TileSizeY
			r.engine.Rectangle(startX, startY, startX+r.opt.TileSizeX, startY+r.opt.TileSizeY, currentTile.BgColor)

			// Once we have the tile background, we need to check if we have an object on this tile.
			// If so, we'll draw the object (Bg & Fg), otherwise we'll draw the tile foreground.
			if objectMap[currentTile.X][currentTile.Y] == 0 {
				// We don't have an entity. Proceed with the foreground
				r.engine.Text(startX, startY, string([]rune{currentTile.Char}), currentTile.FgColor)
			} else {
				// We have an object. First, draw its background (if it's not transparent).
				// TODO: Perform this check *before* rendering the tile background to save a drawing call.
				drawable := objects[objectMap[i][j]].GetComponent(components.DrawableName).(*components.Drawable)
				r.engine.Rectangle(startX, startY, startX+r.opt.TileSizeX, startY+r.opt.TileSizeY, drawable.BgColor)
				r.engine.Text(startX, startY, string([]rune{drawable.Char}), drawable.FgColor)
			}
		}
	}

	// Phase 2 - Commit to screen
	r.engine.Draw()

	return nil
}
