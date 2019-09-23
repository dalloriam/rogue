package cartography

import (
	"math"
)

func heuristic(t1, t2 *Tile) float64 {
	return math.Abs(float64(t1.Position.X())-float64(t2.Position.X())) + math.Abs(float64(t1.Position.Y())-float64(t2.Position.Y()))
}
