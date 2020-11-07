package math3

import "math"

// Sphere represent 3d sphere
type Sphere struct {
	Center Vec
	Radius float64
}

// IsRayIntersected returns true if ray collides with sphere
func (s *Sphere) IsRayIntersected(origin *Vec, direction *Vec, t *float64) bool {
	l := s.Center.Subtract(*origin)
	tca := l.Dot(*direction)
	d2 := l.Dot(l) - tca*tca
	if d2 > s.Radius*s.Radius {
		return false
	}
	thc := math.Sqrt(s.Radius*s.Radius - d2)
	*t = tca - thc
	t1 := tca + thc
	if *t < 0 {
		*t = t1
	}
	if *t < 0 {
		return false
	}
	return true
}
