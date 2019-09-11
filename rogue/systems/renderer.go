package systems

import (
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/entities"
	"image/color"
)

// RenderingBackend abstracts a rendering engine.
type RenderingEngine interface {
	Clear()
	Draw()
	DrawTile(x, y uint64, char rune, fgColor, bgColor color.Color)
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

func (r *Renderer) ShouldTrack(object entities.GameObject) bool {
	return object.HasComponent(components.DrawableName) && object.HasComponent(components.PositionName)
}

// Update updates the system state.
func (r *Renderer) Update(objects map[uint64]entities.GameObject) error {

	r.engine.Clear()

	// TODO: Next iteration of the render process.
	// Phase 1 - Draw map changes (Taking viewport into account?)
	// Phase 2 - Draw all objects.
	// Phase 3 - Commit to screen

	// Update the canvas for all drawable entities.
	for _, obj := range objects {
		drawable := obj.GetComponent(components.DrawableName).(components.Drawable)
		position := obj.GetComponent(components.PositionName).(components.Position)

		r.engine.DrawTile(position.X, position.Y, drawable.Char, drawable.FgColor, drawable.BgColor)
	}

	// Commit the drawing to the screen.
	r.engine.Draw()

	return nil
}
