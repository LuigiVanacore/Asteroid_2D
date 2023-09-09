package asteroid

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"time"
)

type Shoot struct {
	Entity
	aliveTimer *Timer
}

type PlayerShoot struct {
	Shoot
}

func NewShoot(texture *ebiten.Image, world *World, tag int, mask CollisionMask) *Shoot {
	entity := NewEntity(texture, world, tag, mask)
	return &Shoot{Entity: *entity, aliveTimer: NewTimer(time.Second * 5).Start()}
}

func (s *Shoot) Update() {
	s.sprite.Move(s.impulse.X, s.impulse.Y)
	if s.aliveTimer.IsFinished() {
		s.alive = false
	}
}

func NewPlayerShoot(from *Player, texture *ebiten.Image, world *World) *PlayerShoot {
	mask := *NewCollisionMask()
	mask.SetBit(meteor)
	shoot := NewShoot(texture, world, playerShoot, mask)
	shoot.SetPosition(from.GetPosition().X, from.GetPosition().Y)
	shoot.SetRotation(from.GetRotation())
	angle := float64(from.GetRotation())/180*math.Pi - math.Pi/2
	shoot.impulse = Vector2D{math.Cos(angle) * 10, math.Sin(angle) * 10}
	return &PlayerShoot{Shoot: *shoot}
}

func (s *PlayerShoot) IsCollide(other Collidable) bool {
	return s.collider.IsColliding(other)
}

func (s *PlayerShoot) ToString() string {
	return "playerShoot"
}

//
//bool ShootPlayer::isCollide(const Entity& other)const
//{
//if(dynamic_cast<const Enemy*>(&other) != nullptr)
//{
//return Collision::circleTest(_sprite,other._sprite);
//}
//return false;
//}
//
///********************* ShootSaucer *****************/
//
//ShootSaucer::ShootSaucer(SmallSaucer& from) : Shoot(Configuration::Textures::ShootSaucer,from._world)
//{
//_duration = sf::seconds(5);
//
//
//sf::Vector2f pos = Configuration::player->getPosition() - from.getPosition();
//
//float accuracy_lost = book::random(-1.f,1.f)*M_PI/((200+Configuration::getScore())/100.f);
//float angle_rad = std::atan2(pos.y,pos.x) + accuracy_lost;
//float angle_deg = angle_rad * 180 / M_PI;
//
//_impulse = sf::Vector2f(std::cos(angle_rad),std::sin(angle_rad)) * 500.f;
//
//setPosition(from.getPosition());
//_sprite.setRotation(angle_deg + 90);
//_world.add(Configuration::Sounds::LaserEnemy);
//}
//
//bool ShootSaucer::isCollide(const Entity& other)const
//{
//if(dynamic_cast<const Player*>(&other) or dynamic_cast<const Meteor*>(&other))
//{
//return Collision::circleTest(_sprite,other._sprite);
//}
//return false;
//}
//}
