package systems

import (
	"math"

	"github.com/dalloriam/rogue/rogue/structure"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/components"
	"github.com/dalloriam/rogue/rogue/object"
)

const (
	// TODO: Make configurable
	rayCount       = 360
	rayStep        = 3
	darkVisibility = 0.25
)

func makeCameraView(info UpdateInfo) components.CameraView {
	cameraView := make([][]components.ViewInfo, info.WorldMap.SizeX())
	for i := 0; i < info.WorldMap.SizeX(); i++ {
		cameraView[i] = make([]components.ViewInfo, info.WorldMap.SizeY())
		for j := 0; j < info.WorldMap.SizeY(); j++ {
			myViewInfo := components.ViewInfo{Tile: info.WorldMap[i][j]}
			cameraView[i][j] = myViewInfo
		}
	}
	for _, obj := range info.ObjectsByID {
		if !obj.HasComponent(components.PositionName) {
			continue
		}

		position := obj.GetComponent(components.PositionName).(*components.Position)

		// This camera can't see this tile, so we don't register the entities on it in its view.
		if info.WorldMap.At(position).Visibility != 1.0 {
			continue
		}

		objectsAtTile := []object.GameObject{}
		for _, tileObjectID := range info.ObjectPositionMap[position.X()][position.Y()] {
			objectsAtTile = append(objectsAtTile, info.ObjectsByID[tileObjectID])
		}

		cameraView[position.X()][position.Y()] = components.ViewInfo{Tile: *info.WorldMap.At(position), Entities: objectsAtTile}
	}

	return cameraView
}

type SightSystem struct {
	RayCount int
	RayStep  int

	DefaultVisibility cartography.TileVisibility
}

func NewSightSystem(defaultVisibility cartography.TileVisibility) *SightSystem {
	return &SightSystem{
		RayCount:          rayCount,
		RayStep:           rayStep,
		DefaultVisibility: defaultVisibility,
	}
}

func (s *SightSystem) Name() string {
	return "sight"
}

// ShouldTrack returns whether this system should track the object.
func (s *SightSystem) ShouldTrack(object object.GameObject) bool {
	return object.HasComponent(components.PositionName)
}

func (s *SightSystem) Update(info UpdateInfo) error {
	// Make all tiles invisible.
	for i := 0; i < len(info.WorldMap); i++ {
		for j := 0; j < len(info.WorldMap[i]); j++ {
			info.WorldMap.At(structure.V(i, j)).Visibility = s.DefaultVisibility
		}
	}

	for _, obj := range info.ObjectsByID {
		pos := obj.GetComponent(components.PositionName).(*components.Position)

		if !(obj.HasComponent(components.CameraName)) {
			continue
		}

		cam := obj.GetComponent(components.CameraName).(*components.Camera)

		oldView := cam.View
		if len(oldView) == 0 {
			oldView = makeCameraView(info)
		}
		cam.View = makeCameraView(info)

		cam.View.At(pos).Tile.Visibility = 1.0 // Camera always sees its own tile.
		if cam.Main {
			info.WorldMap.At(pos).Visibility = 1.0 // Camera always sees its own tile.
		}

		for i := 0; i < s.RayCount; i += s.RayStep {
			// TODO: Precompute cos values.
			ax := math.Cos(float64(i) / (180.0 / math.Pi))
			ay := math.Sin(float64(i) / (180.0 / math.Pi))

			x := float64(pos.X())
			y := float64(pos.Y())

			for z := 0; z < cam.SightRadius; z++ {
				x += ax
				y += ay

				if x < 0 || y < 0 || int(x) >= info.WorldMap.SizeX() || int(y) >= info.WorldMap.SizeY() {
					break
				}

				// If we reach here, tile {x, y} is visible.
				viewInfo := cam.View.At(structure.V(int(math.Round(x)), int(math.Round(y))))
				viewInfo.Tile.Visibility = 1.0

				if cam.Main {
					info.WorldMap.At(structure.V(int(math.Round(x)), int(math.Round(y)))).Visibility = 1.0
				}

				// However, if the current tile blocks sight, stop raytracing.
				if cam.BlockedBy.Contains(viewInfo.Tile.Type) {
					break
				}
			}
		}

		for i := 0; i < info.WorldMap.SizeX(); i++ {
			for j := 0; j < info.WorldMap.SizeY(); j++ {
				oldViewInfo := oldView.At(structure.V(i, j))
				newViewInfo := cam.View.At(structure.V(i, j))
				if newViewInfo.Tile.Visibility == 0 && oldViewInfo.Tile.Visibility > 0 {
					if cam.Main {
						info.WorldMap[i][j].Visibility = cartography.VisibilityOutOfSight
					}
					newViewInfo.Tile.Visibility = cartography.VisibilityOutOfSight
				}
			}
		}

	}

	// Update observed object position only if object tile is visible by the main camera.
	for _, nonCamObject := range info.ObjectsByID {
		objectPosition := nonCamObject.GetComponent(components.PositionName).(*components.Position)
		if info.WorldMap.At(objectPosition.Vec).Visibility == 1.0 {
			nonCamObject.AddComponents(&components.ObservedPosition{Vec: structure.V(objectPosition.X(), objectPosition.Y())})
		}
	}

	return nil
}
