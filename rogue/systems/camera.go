package systems

import (
	"time"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

type Camera interface {
	MoveTo(x, y int)
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

	for _, obj := range objects {
		focusTgt := obj.GetComponent(components.FocusName).(*components.Focus)
		position := obj.GetComponent(components.PositionName).(*components.Position)

		if focusTgt.Priority > highestPriority {
			bestX = position.X
			bestY = position.Y
		}
	}

	// TODO: Do tile size conversion.
	c.cam.MoveTo(bestX, bestY)

	return nil
}
