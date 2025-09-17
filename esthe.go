package main

import (
	"image/color"
	"log"

	"github.com/hajimeoshi/ebiten/v2"
	"github.com/hajimeoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	x, y      float64
	widht     float64
	height    float64
	ySpeed    float64
	gravity   float64
	jumppower float64
	onGround  bool
}

type player struct {
	x, y   float64
	width  float64
	height float64
}

type Game struct {
	player    Player
	platforms []player
}

func (g *Game) Update() error { // contôle
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.x += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.x -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) && g.player.onGround {
		g.player.ySpeed = -g.player.jumppower
		g.player.onGround = false
	}
	// gravité
	g.player.y += g.player.gravity
	g.player.y += g.player.ySpeed
	g.player.onGround = false

	for _, p := range g.platforms {
		if g.player.y < p.x+p.height > platform.x &&
			g.player.y+g.player.height < platform.y+plat.height &&
			g.player.x+g.player.width > platform.x &&
			g.player.x < platform.x+platform.width {
			g.player.y = platform.y - g.player.height
			g.player.ySpeed = 0
			g.player.onGround = true
		}
	}
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100, 200, 255, 255}) // fond bleu

	// dessiner plateforme
	for _, p := range g.platforms {
		ebitenutil.DrawRect(screen, p.x, p.y, p.width, p.height, color.RGBA{255, 255, 255, 255}) // BLANc

	}
	// dessiner joueur
	ebitenutil.DrawRect(screen, g.player.x, g.player.y, g.player.width, g.player.height, color.RGBA{255, 0, 0, 255}) // rouge
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 400

}
func main() {
	player := Player{
		x: 50, y: 300,
		width: 50, height: 50,
		gravity:   1,
		jumppower: -15,
	}
	platforms := []Player{
		{x: 100, y: 350, width: 200, height: 20},
		{x: 400, y: 300, width: 150, height: 20},
		{x: 600, y: 250, width: 180, height: 20},
	}
	game := &Game{
		player:    player,
		platforms: platforms,
	}
	ebiten.SetWindowSize(800, 400)
	ebiten.SetWindowTitle("Esthe")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
