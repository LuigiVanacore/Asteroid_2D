package asteroid

type Shape interface {
	Intersect(shape Shape) bool
	SetTransform(transform *Transform)
}
