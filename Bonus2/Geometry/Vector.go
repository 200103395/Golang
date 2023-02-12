package Geometry

type Vector struct {
	X int
	Y int
}

func (V *Vector) GetX() int {
	return V.X
}
func (V *Vector) GetY() int {
	return V.Y
}

func (V *Vector) Add(vector Vector) *Vector {
	newVector := &Vector{X: V.X + vector.X, Y: V.Y + vector.Y}
	return newVector
}
