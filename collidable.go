package asteroid

type Collidable interface {
	Tagable
	IsCollide(other Collidable) bool
	GetShape() Shape
	ToString() string
}
