package main

import (
	asteroid "Asteroid_2D"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowSize(asteroid.WindowWidth, asteroid.WindowHeight)
	ebiten.SetWindowTitle("AirWars2D")
	game := &asteroid.Game{}
	asteroid.SetDebug(true)
	game.Init()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
