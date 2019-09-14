package generator

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/dalloriam/rogue/rogue/cartography"
)

type Rectangle struct {
	// TODO: Extract in some kind of "map building kit" package.
	StartX int
	EndX   int
	StartY int
	EndY   int

	CenterX int
	CenterY int
}

func NewRectangle(x, y, w, h float64) Rectangle {
	r := Rectangle{
		StartX: int(x),
		EndX:   int(x + w),
		StartY: int(y),
		EndY:   int(y + h),
	}
	r.CenterX = int(math.Round(float64(r.StartX + r.EndX/2)))
	r.CenterY = int(math.Round(float64(r.StartY + r.EndY/2)))

	return r
}

// Intersects computes the intersection between this rectangle and another.
func (r Rectangle) Intersects(other Rectangle) bool {
	return r.StartX <= other.EndX && r.EndX >= other.StartX && r.StartY <= other.EndY && r.EndY >= other.StartY
}

// DungeonGenerator generates a dungeon.
type DungeonGenerator struct {
	MaxRoomSize      int
	MinRoomSize      int
	MaxNumberOfRooms int

	levelMap cartography.Map
}

func NewDungeonGenerator(maxRoomSize, minRoomSize, maxNumberOfRooms, mapSizeX, mapSizeY int) *DungeonGenerator {
	return &DungeonGenerator{
		MaxRoomSize:      maxRoomSize,
		MinRoomSize:      minRoomSize,
		MaxNumberOfRooms: maxNumberOfRooms,
		levelMap:         cartography.NewMap(mapSizeX, mapSizeY),
	}
}

func (g *DungeonGenerator) fillMapWithRockWalls() {
	for i := 0; i < g.levelMap.SizeX(); i++ {
		for j := 0; j < g.levelMap.SizeY(); j++ {
			g.levelMap.Set(i, j, cartography.Tile{
				X:       i,
				Y:       j,
				Char:    '#',
				FgColor: color.Black,
				BgColor: color.Gray{128},
			})
		}
	}
}

func (g *DungeonGenerator) digRectangle(r Rectangle) {
	for i := r.StartX; i < r.EndX; i++ {
		for j := r.StartY; j < r.EndY; j++ {
			g.levelMap.Set(i, j, cartography.Tile{
				X:       i,
				Y:       j,
				Char:    '.',
				FgColor: color.White,
				BgColor: color.Black,
			})
		}
	}
}

func (g *DungeonGenerator) digVerticalTunnel(startY, endY, x int) {
	for y := math.Min(float64(startY), float64(endY)); y < math.Max(float64(startY), float64(endY)); y++ {
		g.levelMap.Set(x, int(y), cartography.Tile{
			X:       x,
			Y:       int(y),
			Char:    '.',
			FgColor: color.White,
			BgColor: color.Black,
		})
	}
}
func (g *DungeonGenerator) digHorizontalTunnel(startX, endX, y int) {
	fmt.Println(startX, endX)
	for x := math.Min(float64(startX), float64(endX)); x < math.Max(float64(startX), float64(endX)); x++ {
		g.levelMap.Set(int(x), y, cartography.Tile{
			X:       int(x),
			Y:       y,
			Char:    '.',
			FgColor: color.White,
			BgColor: color.Black,
		})
	}
}

func (g *DungeonGenerator) Generate() cartography.Map {
	// Initialize an empty map.
	g.fillMapWithRockWalls()

	numRooms := 0
	var rooms []Rectangle

	for i := 0; i < g.MaxNumberOfRooms; i++ {
		w := math.Floor(rand.Float64()*float64(g.MaxRoomSize-g.MinRoomSize-1)) + float64(g.MinRoomSize)
		h := math.Floor(rand.Float64()*float64(g.MaxRoomSize-g.MinRoomSize-1)) + float64(g.MinRoomSize)

		x := math.Floor(rand.Float64()*(float64(g.levelMap.SizeX())-w-2)) + 1
		y := math.Floor(rand.Float64()*(float64(g.levelMap.SizeY())-h-2)) + 1

		failing := true
		var newRoom Rectangle

		for failing {
			newRoom = NewRectangle(x, y, w, h)

			// Make sure this room intersects with no other room.
			failing = false
			for _, otherRoom := range rooms {
				if otherRoom.Intersects(newRoom) {
					failing = true
				}
			}
		}

		g.digRectangle(newRoom)

		// Once the new room is added, dig a tunnel from the previous room to make it accessible.
		if len(rooms) > 0 {
			prevX := rooms[numRooms-1].CenterX
			prevY := rooms[numRooms-1].CenterY

			fmt.Println("HERE")
			if rand.Float64() > 0.5 {
				g.digHorizontalTunnel(prevX, newRoom.CenterX, prevY)
				g.digVerticalTunnel(prevY, newRoom.CenterY, newRoom.CenterX)
			} else {
				g.digVerticalTunnel(prevY, newRoom.CenterY, prevX)
				g.digHorizontalTunnel(prevX, newRoom.CenterX, newRoom.CenterY)
			}
		}

		rooms = append(rooms, newRoom)
		numRooms++
	}

	return g.levelMap
}
