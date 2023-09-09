package asteroid

type Transformable interface {
	GetTransform() *Transform
	SetTransform(transform Transform)
}
