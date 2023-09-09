package asteroid

import (
	"Asteroid_2D/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

const (
	WindowWidth  = 800
	WindowHeight = 640
)

var debug bool

func SetDebug(isDebug bool) {
	debug = isDebug
}

const MAX_PLAYER_LIFES int = 3

type Game struct {
	ready       bool
	playerLives int
	world       *World
	player      *Player
}

func (g *Game) Init() {
	g.world = NewWorld(WindowWidth, WindowHeight)
	g.PlayerRespawn()
	g.initLevel()
}

func (g *Game) initLevel() {
	nb_meteors := 1
	g.playerLives = MAX_PLAYER_LIFES

	for i := 0; i < nb_meteors; i++ {
		meteor := NewBigMeteor(g.world)
		for g.world.IsCollide(meteor) {
			meteor.SetPosition(rand.Float64()*WindowWidth, rand.Float64()*WindowHeight)
		}
		g.world.Add(meteor)
	}
}

func (g *Game) HandleInput() error {

	return nil
}

func (g *Game) GameOver() {
	g.world.Clear()
}

func (g *Game) PlayerRespawn() {
	player := NewPlayer(assets.ResourceManager().GetTexture(assets.Ship), g.world)
	player.SetPosition(50, 50)
	g.world.Add(player)
	g.player = player
}

func (g *Game) Update() error {
	g.world.Update()
	if !g.player.alive {
		g.playerLives--
		if g.playerLives <= 0 {
			g.GameOver()
		} else {
			g.PlayerRespawn()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	g.world.Draw(screen, op)
}

func (g *Game) Layout(int, int) (screenwidth int, screenheight int) {
	return WindowWidth, WindowHeight
}
