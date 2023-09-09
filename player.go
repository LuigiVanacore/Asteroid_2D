package asteroid

import (
	"Asteroid_2D/assets"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"time"
)

type direction uint

const (
	Left direction = iota
	Right
	Up
)

const (
	moveLeft = iota
	moveRight
	moveUp
	playerShoot
)

type Player struct {
	Entity
	*ActionTarget
	shootTimer     *Timer
	invicibleTimer *Timer
	canMove        bool
}

func NewPlayer(texture *ebiten.Image, world *World) *Player {
	mask := *NewCollisionMask()
	mask.SetBit(meteor)
	player := &Player{Entity: *NewEntity(texture, world, player, mask), shootTimer: NewTimer(300 * time.Millisecond), invicibleTimer: NewTimer(100 * time.Millisecond)}
	player.setUpInputs()
	return player
}

func (p *Player) setUpInputs() {
	actionMap := NewActionMap()
	actionLeft := NewActionKey(ebiten.KeyLeft, PRESSED)
	actionRight := NewActionKey(ebiten.KeyRight, PRESSED)
	actionUp := NewActionKey(ebiten.KeyUp, PRESSED)
	actionShoot := NewActionKey(ebiten.KeySpace, PRESSED)
	actionMap.Add(moveLeft, *actionLeft)
	actionMap.Add(moveRight, *actionRight)
	actionMap.Add(moveUp, *actionUp)
	actionMap.Add(playerShoot, *actionShoot)
	actionTarget := NewActionTarget(actionMap)
	actionTarget.Bind(moveLeft, func() { p.canMove = true; p.Move(Left) })
	actionTarget.Bind(moveRight, func() { p.canMove = true; p.Move(Right) })
	actionTarget.Bind(moveUp, func() {
		p.canMove = true
		p.Move(Up)
	})
	actionTarget.Bind(playerShoot, func() {
		p.shoot()
	})
	p.ActionTarget = actionTarget
}

func (p *Player) Update() {
	p.ProcessEvents()
}

func (p *Player) ProcessEvents() {
	p.canMove = false
	p.ActionTarget.ProcessEvents()
}

func (p *Player) Move(direction direction) {
	if p.canMove {
		if direction == Left {
			p.GetTransform().Rotate(-5)
		}
		if direction == Right {
			p.GetTransform().Rotate(5)
		}
		if direction == Up {
			angle := float64(p.GetTransform().GetRotation())/180*math.Pi - math.Pi/2
			p.GetTransform().Move(math.Cos(angle)*5, math.Sin(angle)*5)
		}
	}
}

func (p *Player) ToString() string {
	return "player"
}

func (p *Player) shoot() {

	if p.shootTimer.IsFinished() {
		shoot := NewPlayerShoot(p, assets.ResourceManager().GetTexture(playerShoot), p.world)
		p.world.Add(shoot)
		fmt.Println("shoot!")
	}
}

func (p *Player) OnDestroy() {
	p.alive = false
}
