package systems

// RenderingBackend abstracts a rendering engine.
type RenderingBackend interface {
	DrawTile()
}

// A Renderer renders components.
type Renderer struct {
}

// NewRenderer returns a new rendering system.
func NewRenderer() (*Renderer, error) {
	return &Renderer{}, nil
}

// Update updates the system state.
func (r *Renderer) Update() {

}
