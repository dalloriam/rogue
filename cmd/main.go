package main

import (
	"image/color"
	"time"

	"github.com/dalloriam/rogue/rogue/cartography"

	"github.com/dalloriam/rogue/cmd/generator"

	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"

	"github.com/dalloriam/rogue/rogue/systems"

	"github.com/dalloriam/rogue/rogue"

	"github.com/dalloriam/rogue/backends/roguepixel"

	"github.com/faiface/pixel/pixelgl"
)

func findPlayer(level cartography.Map) (int, int) {
	// Locate player coordinates
	// TODO: Improve randomness of player position.
	for i := 0; i < level.SizeX(); i++ {
		for j := 0; j < level.SizeY(); j++ {
			if level.At(i, j).Type == "floor" {
				return i, j
			}
		}
	}
	panic("no suitable space")
}

func pixelRun() {

	// Rogue renderer setup.
	opt := roguepixel.GridRenderOptions{
		FontFacePath: "data/font.ttf",
		FontSize:     19,

		TileHeight: 25,
		TileWidth:  25,

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
	renderingSystem, err := systems.NewRenderer(r)
	if err != nil {
		panic(err)
	}

	// Creating the world.
	world := rogue.NewWorld()
	world.AddSystem(renderingSystem, 1)
	world.AddSystem(systems.NewMovementSystem(), 2)
	world.AddSystem(systems.NewControllerSystem(roguepixel.NewInputHandler(r.Window)), 999)

	gen := generator.NewDungeonGenerator(
		10,
		6,
		20,
		int(float64(opt.WindowSizeX)/float64(opt.TileWidth)), int(float64(opt.WindowSizeY)/float64(opt.TileHeight)),
	)

	lvlManager := cartography.NewLevelManager("test.txt", time.Now().UnixNano())
	lvlManager.AddLevel("dungeon_1", gen)
	lvl, ok := lvlManager.GetLevel("dungeon_1")
	if !ok {
		panic("level does not exist")
	}
	world.LoadMap(lvl)

	playerX, playerY := findPlayer(lvl)

	player := object.New(
		&components.Drawable{
			Char:    '@',
			FgColor: color.White,
			BgColor: color.RGBA{0, 0, 0, 0},
		},
		&components.Position{
			X: playerX,
			Y: playerY,
		},
		&components.Physics{BlockedBy: []string{"wall"}},
		&components.PlayerControl{},
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
