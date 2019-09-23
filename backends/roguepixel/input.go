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

// RepeatModeTriggered returns if repeat mode was triggered.
func (i *InputHandler) RepeatModeTriggered() bool {
	return i.window.JustPressed(pixelgl.KeyW)
}

func (i *InputHandler) AutoExploreTriggered() bool {
	return i.window.JustPressed(pixelgl.KeyX)
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

	if i.window.JustPressed(pixelgl.KeyN) {
		return cartography.DirectionDownRight
	}

	if i.window.JustPressed(pixelgl.KeyB) {
		return cartography.DirectionDownLeft
	}

	if i.window.JustPressed(pixelgl.KeyY) {
		return cartography.DirectionUpLeft
	}

	if i.window.JustPressed(pixelgl.KeyU) {
		return cartography.DirectionUpRight
	}

	return cartography.NoDirection
}
