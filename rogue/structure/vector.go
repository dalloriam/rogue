package structure

type Vec interface {
	X() int
	Y() int

	Add(other Vec)
}

type vec struct {
	x int
	y int
}

func V(x, y int) Vec {
	return &vec{x, y}
}

func (v *vec) X() int {
	return v.x
}

func (v *vec) Y() int {
	return v.y
}

func (v *vec) Add(other Vec) {
	v.x += other.X()
	v.y += other.Y()
}
