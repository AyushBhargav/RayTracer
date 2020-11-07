package math3

// Vec denotes vector of 3 dimensions
type Vec struct {
	X float64
	Y float64
	Z float64
}

// Get returns dimension of Vector based on index
func (vec *Vec) Get(index int) float64 {
	switch index {
	case 0:
		return vec.X
	case 1:
		return vec.Y
	case 2:
		return vec.Z
	}

	panic("Invalid index: " + string(index))
}

// Subtract returns new vector with vec - vec2
func (vec *Vec) Subtract(vec2 Vec) Vec {
	return Vec{vec.X - vec2.X, vec.Y - vec2.Y, vec.Z - vec2.Z}
}

// Multiply returns new vector with vec*t
func (vec *Vec) Multiply(t float64) Vec {
	return Vec{vec.X * t, vec.Y * t, vec.Z * t}
}

// Dot returns dot product of vec and vec2
func (vec *Vec) Dot(vec2 Vec) float64 {
	return vec.X*vec2.X + vec.Y*vec2.Y + vec.Z*vec2.Z
}
