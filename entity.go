package asteroid

import "github.com/hajimehoshi/ebiten/v2"

const (
	player = iota
	meteor
	shoot
)

type Entity struct {
	collider *CollisionShape
	impulse  Vector2D
	alive    bool
	sprite   *Sprite
	world    *World
	tag      int
}

func (e *Entity) Update() {
}

func NewEntity(texture *ebiten.Image, world *World, tag int, mask CollisionMask) *Entity {
	sprite := NewSprite(texture)
	sprite.SetPivotToCenter()
	circle := NewCircle(40)
	collider := NewCollisionShape(sprite.GetTransform(), circle, mask)
	return &Entity{sprite: sprite, world: world, alive: true, tag: tag, collider: collider}
}

func (e *Entity) IsCollide(other Collidable) bool {
	return e.collider.IsColliding(other)
}

func (e *Entity) GetShape() Shape {
	return e.collider.shape
}

func (e *Entity) GetTag() int {
	return e.tag
}

func (e *Entity) IsAlive() bool {
	return e.alive
}

func (e *Entity) GetPosition() Vector2D {
	return e.sprite.GetPosition()
}

func (e *Entity) SetPosition(x, y float64) {
	e.sprite.SetPosition(x, y)
}

func (e *Entity) SetRotation(rotation int) {
	e.sprite.rotation = rotation
}

func (e *Entity) GetRotation() int {
	return e.sprite.rotation
}

func (e *Entity) GetPivot() Vector2D {
	return e.sprite.pivot
}

func (e *Entity) GetTransform() *Transform {
	return e.sprite.GetTransform()
}

func (e *Entity) SetTransform(transform Transform) {
	e.sprite.SetTransform(transform)
}

func (e *Entity) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	//op.GeoM = e.geoM
	e.sprite.Draw(target, op)
	if debug {
		e.collider.DrawDebug(target, op)
	}
}

func (e *Entity) OnDestroy() {
	e.alive = false
}
