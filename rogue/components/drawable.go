package components

import "image/color"

// Name of the component
const (
	DrawableName = "drawable"
)

// Drawable represents a drawable.
type Drawable struct {
	Char    rune
	FgColor color.Color
	BgColor color.Color
}

// Name returns the component's name
func (d *Drawable) Name() string {
	return DrawableName
}
