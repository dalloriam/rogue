package main

import (
	"image/color"

	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/entities"

	"github.com/dalloriam/rogue/rogue/cartography"
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
	renderingSystem, err := systems.NewRenderer(r, systems.RendererOptions{
		TileSizeX: 30,
		TileSizeY: 30,
	})
	if err != nil {
		panic(err)
	}

	// Creating the world.
	world := rogue.NewWorld()
	world.AddSystem(renderingSystem, 1)

	worldMap := cartography.NewMap(20, 20)

	var i, j uint64
	for i = 0; i < 20; i++ {
		for j = 0; j < 20; j++ {
			worldMap.Set(i, j, cartography.Tile{
				X:       i,
				Y:       j,
				Char:    '#',
				FgColor: color.White,
				BgColor: color.RGBA{128, 0, 0, 255},
			})
		}
	}
	world.LoadMap(worldMap)

	player := entities.NewObject(
		components.Drawable{
			Char:    '@',
			FgColor: color.White,
			BgColor: color.RGBA{0, 0, 0, 0},
		},
		components.Position{
			X: 10,
			Y: 10,
		},
	)
	world.AddObject(player)

	for r.Running() {
		if err := world.Tick(); err != nil {
			panic(err)
		}
	}
}

func main() {
	pixelgl.Run(pixelRun)
}
