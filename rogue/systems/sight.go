package systems

import (
	"math"
	"time"

	"github.com/dalloriam/rogue/rogue/structure"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

type SightSystem struct {
	RayCount int
	RayStep  int
}

func NewSightSystem() *SightSystem {
	// TODO: Extract constants.
	return &SightSystem{
		RayCount: 360,
		RayStep:  3,
	}
}

// ShouldTrack returns whether this system should track the object.
func (s *SightSystem) ShouldTrack(object object.GameObject) bool {
	return object.HasComponent(components.CameraName) && object.HasComponent(components.PositionName)
}

func (s *SightSystem) Update(dT time.Duration, worldMap cartography.Map, objects map[uint64]object.GameObject) error {
	// Make all tiles invisible.
	for i := 0; i < len(worldMap); i++ {
		for j := 0; j < len(worldMap[i]); j++ {
			worldMap.At(structure.V(i, j)).Visibility = 0.0
		}
	}

	for _, obj := range objects {
		pos := obj.GetComponent(components.PositionName).(*components.Position)
		cam := obj.GetComponent(components.CameraName).(*components.Camera)

		worldMap.At(pos).Visibility = 1.0 // Camera always sees its own tile.

		for i := 0; i < s.RayCount; i += s.RayStep {
			ax := math.Cos(float64(i) / (180.0 / math.Pi))
			ay := math.Sin(float64(i) / (180.0 / math.Pi))

			x := float64(pos.X())
			y := float64(pos.Y())

			for z := 0; z < cam.SightRadius; z++ {
				x += ax
				y += ay

				if x < 0 || y < 0 || int(x) >= worldMap.SizeX() || int(y) >= worldMap.SizeY() {
					break
				}

				// If we reach here, tile {x, y} is visible.
				tile := worldMap.At(structure.V(int(math.Round(x)), int(math.Round(y))))
				tile.Visibility = 1.0

				cam.Memory = append(cam.Memory, tile.Position)

				// However, if the current tile blocks sight, stop raytracing.
				if cam.BlockedBy.Contains(tile.Type) {
					break
				}
			}
		}

		// Override tile memory
		for _, tileVec := range cam.Memory {
			if t := worldMap.At(tileVec); t.Visibility == 0.0 {
				t.Visibility = 0.25 // TODO: Store somewhere
			}
		}
	}
	return nil
}
