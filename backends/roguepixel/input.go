package roguepixel

import (
	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/faiface/pixel/pixelgl"
)

// InputHandler handles input using pixel.
type InputHandler struct {
	window *pixelgl.Window
}

func NewInputHandler(win *pixelgl.Window) *InputHandler {
	return &InputHandler{
		window: win,
	}
}

func (i *InputHandler) GetDirection() cartography.Direction {
	if i.window.JustPressed(pixelgl.KeyK) {
		return cartography.DirectionUp
	}

	if i.window.JustPressed(pixelgl.KeyJ) {
		return cartography.DirectionDown
	}

	if i.window.JustPressed(pixelgl.KeyL) {
		return cartography.DirectionRight
	}

	if i.window.JustPressed(pixelgl.KeyH) {
		return cartography.DirectionLeft
	}

	return cartography.NoDirection
}
