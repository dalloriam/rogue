package main

import (
	"image/color"

	"github.com/dalloriam/rogue/rogue/entities"
	"github.com/dalloriam/rogue/rogue/systems"

	"github.com/dalloriam/rogue/rogue"

	"github.com/dalloriam/rogue/backends/roguepixel"

	"github.com/faiface/pixel/pixelgl"
)

func pixelRun() {

	// Rogue renderer setup.
	r, err := roguepixel.NewRenderer(roguepixel.GridRenderOptions{
		FontFacePath: "data/font.ttf",
		FontSize:     22,
		TileSizeX:    30,
		TileSizeY:    30,

		WindowTitle: "Rogue Demo",
		WindowSizeX: 1024,
		WindowSizeY: 768,

		SmoothDrawing: true,
		VSync:         true,
	})
	if err != nil {
		panic(err)
	}

	// Creating system from rogue renderer.
	renderingSystem, err := systems.NewRenderer(r)
	if err != nil {
		panic(err)
	}

	// Creating the world.
	world := rogue.NewWorld()
	world.AddSystem(renderingSystem, 1)

	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			world.AddObject(entities.NewTile(uint64(i), uint64(j), '#', color.White, color.RGBA{128, 0, 0, 255}))
		}
	}

	for r.Running() {
		if err := world.Tick(); err != nil {
			panic(err)
		}
	}
}

func main() {
	pixelgl.Run(pixelRun)
}