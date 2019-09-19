package structure

type Vec struct {
	X int
	Y int
}

func V(x, y int) Vec {
	return Vec{x, y}
}
