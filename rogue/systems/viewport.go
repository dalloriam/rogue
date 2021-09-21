package systems

import (
	"github.com/dalloriam/rogue/rogue/structure"

	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

type Viewport interface {
	Move(x, y int)
	SetZoom(amount float64)
}

type ViewportSystem struct {
	cam Viewport
}

func NewViewportSystem(camera Viewport) *ViewportSystem {
	return &ViewportSystem{cam: camera}
}

func (v *ViewportSystem) Name() string {
	return "viewport"
}

func (c *ViewportSystem) ShouldTrack(object object.GameObject) bool {
	return object.HasComponent(components.FocusName) && object.HasComponent(components.PositionName)
}

func (c *ViewportSystem) Update(info UpdateInfo) error {
	var bestPos structure.Vec
	highestPriority := -1
	punctual := false
	var bestObject object.GameObject

	for _, obj := range info.ObjectsByID {
		focusTgt := obj.GetComponent(components.FocusName).(*components.Focus)
		position := obj.GetComponent(components.PositionName).(*components.Position)

		if focusTgt.Priority > highestPriority {
			bestPos = position
			bestObject = obj
			punctual = focusTgt.Punctual
		}
	}

	if bestPos != nil {
		// We only move the camera if one exists.
		c.cam.Move(bestPos.X(), bestPos.Y())

		if punctual && bestObject != nil {
			bestObject.RemoveComponent(components.FocusName)
		}
	}

	return nil
}
