package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

func drawRectangle(imd *imdraw.IMDraw, origin pixel.Vec, sizeX float64, sizeY float64, color pixel.RGBA) {
	imd.Color = color
	imd.Push(origin)

	origin.X += sizeX
	origin.Y += sizeY

	imd.Push(origin)

	imd.Rectangle(0)
}

func pixelRun() {
	cfg := pixelgl.WindowConfig{
		Title:  "Rogue",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	drawRectangle(imd, pixel.V(400, 400), 100, 300, pixel.RGB(1, 0, 0))
	drawRectangle(imd, pixel.V(500, 400), 100, 300, pixel.RGB(0, 1, 0))
	drawRectangle(imd, pixel.V(600, 400), 100, 300, pixel.RGB(0, 0, 1))

	for !win.Closed() {
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(pixelRun)
}
