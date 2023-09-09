package asteroid

import (
	"Asteroid_2D/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type Meteor struct {
	Enemy
}

func NewMeteor(texture *ebiten.Image, world *World) *Meteor {
	mask := *NewCollisionMask()
	mask.SetBit(player)
	mask.SetBit(playerShoot)
	return &Meteor{Enemy: *NewEnemy(texture, world, meteor, mask)}
}

func (m *Meteor) IsCollide(other Collidable) bool {
	return m.collider.IsColliding(other)
}

func (m *Meteor) Update() {
	m.sprite.Move(m.impulse.X, m.impulse.Y)
}

type BigMeteor struct {
	Meteor
}

func NewBigMeteor(world *World) *BigMeteor {
	meteor := NewMeteor(assets.ResourceManager().GetTexture(assets.BigMeteor), world)
	meteor.SetPoints(20)
	meteor.impulse = meteor.impulse.MultiplyScalar(2)
	return &BigMeteor{Meteor: *meteor}
}

func (s *BigMeteor) ToString() string {
	return "bigMeteor"
}

func (b *BigMeteor) OnDestroy() {
	b.Meteor.Entity.OnDestroy()

	numDebris := rand.Intn(2) + 2
	for i := 0; i < numDebris; i++ {
		mediumMeteor := NewBigMeteor(b.world)
		mediumMeteor.SetPosition(b.GetPosition().X, b.GetPosition().Y)
		b.world.Add(mediumMeteor)
	}
}

type MediumMeteor struct {
	Meteor
}

func NewMediumMeteor(world *World) *MediumMeteor {
	meteor := NewMeteor(nil, world)
	meteor.SetPoints(60)
	meteor.impulse = meteor.impulse.MultiplyScalar(200)
	return &MediumMeteor{Meteor: *meteor}
}
