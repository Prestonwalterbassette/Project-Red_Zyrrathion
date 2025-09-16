package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	image *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screenWidth, screenHeight := screen.Size()
	imgWidth := g.image.Bounds().Dx()
	imgHeight := g.image.Bounds().Dy()

	scaleX := float64(screenWidth) / float64(imgWidth)
	scaleY := float64(screenHeight) / float64(imgHeight)
	scale := scaleX
	if scaleY < scaleX {
		scale = scaleY
	}

	x := (float64(screenWidth) - float64(imgWidth)*scale) / 2
	y := (float64(screenHeight) - float64(imgHeight)*scale) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x, y)

	screen.DrawImage(g.image, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	img, _, err := ebitenutil.NewImageFromFile("100034886.png")
	if err != nil {
		log.Fatal(err) // ⬅️ Vérification immédiate
	}

	game := &Game{
		image: img,
	}

	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Image Plein Écran")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
