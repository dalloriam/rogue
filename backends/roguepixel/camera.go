package roguepixel

import (
	"github.com/faiface/pixel"
)

type Camera struct {
	Position pixel.Vec
	Zoom     float64

	tileWidth  int
	tileHeight int

	leftBound, rightBound, bottomBound, topBound float64
}

func NewCamera(tileWidth, tileHeight, mapWidth, mapHeight, viewportWidth, viewportHeight int) *Camera {

	leftBound := viewportWidth / 2
	rightBound := (mapWidth * tileWidth) - leftBound

	bottomBound := viewportHeight / 2
	topBound := (mapHeight * tileHeight) - bottomBound

	return &Camera{
		Position: pixel.ZV,
		Zoom:     1.0,

		tileWidth:   tileWidth,
		tileHeight:  tileHeight,
		leftBound:   float64(leftBound),
		rightBound:  float64(rightBound),
		topBound:    float64(topBound),
		bottomBound: float64(bottomBound),
	}
}

// Zoom zooms the camera. 1.0 is 100%.
func (c *Camera) SetZoom(amount float64) {
	c.Zoom = amount
}

func (c *Camera) Move(x, y int) {
	actX := float64(x * c.tileWidth)
	if actX < c.leftBound {
		actX = c.leftBound
	} else if actX > c.rightBound {
		actX = c.rightBound
	}

	actY := float64(y * c.tileHeight)
	if actY < c.bottomBound {
		actY = c.bottomBound
	} else if actY > c.topBound {
		actY = c.topBound
	}

	c.Position.X = actX
	c.Position.Y = actY
}
