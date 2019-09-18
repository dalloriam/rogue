package systems

import (
	"time"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

type Camera interface {
	Move(x, y int)
	SetZoom(amount float64)
}

type CameraSystem struct {
	cam Camera
}

func NewCameraSystem(camera Camera) *CameraSystem {
	return &CameraSystem{cam: camera}
}

func (c *CameraSystem) ShouldTrack(object object.GameObject) bool {
	return object.HasComponent(components.FocusName) && object.HasComponent(components.PositionName)
}

func (c *CameraSystem) Update(dT time.Duration, worldMap cartography.Map, objects map[uint64]object.GameObject) error {
	var bestX, bestY int
	highestPriority := -1
	punctual := false
	var bestObject object.GameObject

	for _, obj := range objects {
		focusTgt := obj.GetComponent(components.FocusName).(*components.Focus)
		position := obj.GetComponent(components.PositionName).(*components.Position)

		if focusTgt.Priority > highestPriority {
			bestX = position.X
			bestY = position.Y
			bestObject = obj
			punctual = focusTgt.Punctual
		}
	}

	// TODO: Do tile size conversion.
	c.cam.Move(bestX, bestY)

	if punctual && bestObject != nil {
		bestObject.RemoveComponent(components.FocusName)
	}

	return nil
}
