package ui

import (
	"image/color"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Bouton représente un bouton cliquable
type Bouton struct {
	X, Y          int
	Width, Height int
	Label         string
	Color         color.Color
	Font          font.Face
}
type FenetreGame struct {
	background *ebiten.Image
	bouton     *Bouton
	mouseDown  bool
	clicked    bool
}

func LoadFont(path string, size float64) font.Face {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	tt, err := opentype.Parse(data)
	if err != nil {
		log.Fatal(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	return face
}
func ColorRGB(r, g, b int) color.Color {
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

func NewFenetre(background *ebiten.Image, bouton *Bouton) *FenetreGame {
	return &FenetreGame{
		background: background,
		bouton:     bouton,
	}
}

// Update gère le clic sur le bouton
func (g *FenetreGame) Update() error {
	x, y := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !g.mouseDown {
			g.mouseDown = true
			if g.bouton.IsClicked(x, y) {
				g.clicked = true
				return ErrQuit
			}
		}
	} else {
		g.mouseDown = false
	}
	return nil
}

// Draw dessine le bouton et le fond
func (g *FenetreGame) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.background, nil)
	if !g.clicked {
		g.bouton.Draw(screen)
	}
}

// Layout fixe la taille de la fenêtre
func (g *FenetreGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1920, 1080
}

// Draw du bouton
func (b *Bouton) Draw(screen *ebiten.Image) {
	rect := ebiten.NewImage(b.Width, b.Height)
	rect.Fill(b.Color)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.X), float64(b.Y))
	screen.DrawImage(rect, op)

	text.Draw(screen, b.Label, b.Font, b.X+10, b.Y+b.Height-10, color.White)
}

// Vérifie si le bouton est cliqué
func (b *Bouton) IsClicked(mouseX, mouseY int) bool {
	return mouseX >= b.X && mouseX <= b.X+b.Width &&
		mouseY >= b.Y && mouseY <= b.Y+b.Height
}

var ErrQuit = &QuitError{}

type QuitError struct{}

func (e *QuitError) Error() string { return "fenêtre terminée" }

// LancerFenetre ouvre la fenêtre et retourne dès que "Jouer" est cliqué
func LancerFenetre(background *ebiten.Image, bouton *Bouton) {
	game := NewFenetre(background, bouton)
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Zyrrathion")
	err := ebiten.RunGame(game)
	if err != nil && err != ErrQuit {
		log.Fatal(err)
	}
}
