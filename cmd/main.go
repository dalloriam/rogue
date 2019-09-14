package main

import (
	"image/color"

	"github.com/dalloriam/rogue/cmd/generator"

	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/entities"

	"github.com/dalloriam/rogue/rogue/systems"

	"github.com/dalloriam/rogue/rogue"

	"github.com/dalloriam/rogue/backends/roguepixel"

	"github.com/faiface/pixel/pixelgl"
)

func pixelRun() {

	// Rogue renderer setup.
	opt := roguepixel.GridRenderOptions{
		FontFacePath: "data/font.ttf",
		FontSize:     19,

		WindowTitle: "Rogue Demo",
		WindowSizeX: 1300,
		WindowSizeY: 900,

		SmoothDrawing: true,
		VSync:         true,
	}
	r, err := roguepixel.NewRenderer(opt)
	if err != nil {
		panic(err)
	}

	// Creating system from rogue renderer.
	renderOpt := systems.RendererOptions{
		TileSizeX: 25,
		TileSizeY: 25,
	}
	renderingSystem, err := systems.NewRenderer(r, renderOpt)
	if err != nil {
		panic(err)
	}

	// Creating the world.
	world := rogue.NewWorld()
	world.AddSystem(renderingSystem, 1)

	gen := generator.NewDungeonGenerator(
		10,
		6,
		3,
		int(float64(opt.WindowSizeX)/float64(renderOpt.TileSizeX)), int(float64(opt.WindowSizeY)/float64(renderOpt.TileSizeY)),
	)
	world.LoadMap(gen.Generate())

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
