package asteroid

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"math/rand"
)

type Enemy struct {
	Entity
	points int
}

func NewEnemy(texture *ebiten.Image, world *World, tag int, mask CollisionMask) *Enemy {
	entity := NewEntity(texture, world, tag, mask)
	angle := rand.Float64() * 2 * math.Pi
	entity.impulse = Vector2D{math.Cos(angle), math.Sin(angle)}
	return &Enemy{Entity: *entity}
}

func (e *Enemy) getPoints() int {
	return e.points
}

func (e *Enemy) ToString() string {
	return "enemy"
}

func (e *Enemy) OnDestroy() {
	e.Entity.OnDestroy()
	//e.world.AddScore(e.getPoints())
}

func (e *Enemy) SetPoints(points int) {
	e.points = points
}
