package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var myFont font.Face

func init() {
	// Charger la police (ex: Arial.ttf)
	b, err := os.ReadFile("Arial.ttf")
	if err != nil {
		log.Fatal(err)
	}
	tt, err := opentype.Parse(b)
	if err != nil {
		log.Fatal(err)
	}

	// Créer un face avec la taille souhaitée
	myFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48, // taille du texte
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func drawText(screen *ebiten.Image, x, y int, msg string, col color.Color) {
	text.Draw(screen, msg, myFont, x, y, col)
}

type Game struct {
	background *ebiten.Image
	mouseDown  bool
}

func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !g.mouseDown {
			g.mouseDown = true
			// Hitbox du bouton "Jouer"
			if x >= 920 && x <= 920+200 && y >= 900-48 && y <= 900 { // 48 = taille du texte
				fmt.Println("Bouton 'Jouer' cliqué !")
			}
		}
	} else {
		g.mouseDown = false
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Redimensionner l’image au format fenêtre
	op := &ebiten.DrawImageOptions{}
	sw, sh := g.background.Size()
	scaleX := float64(1920) / float64(sw)
	scaleY := float64(1080) / float64(sh)
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(g.background, op)

	// Bouton rouge avec texte agrandi
	drawText(screen, 920, 900, "Jouer", color.White)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1920, 1080
}

func main() {
	// Charger l'image de fond
	file, err := os.Open("Zyrrathion.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	game := &Game{
		background: ebiten.NewImageFromImage(img),
		mouseDown:  false,
	}

	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Zyrrathion")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}