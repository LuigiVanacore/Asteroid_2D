package asteroid

type Actor interface {
	Tagable
	Collidable
	Transformable
	Drawable
	Updatable
	IsAlive() bool
	OnDestroy()
}
