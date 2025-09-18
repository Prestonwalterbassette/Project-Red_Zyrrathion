package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

// Bouton représente un bouton cliquable
type Bouton struct {
	X, Y          int
	Width, Height int
	Label         string
	Color         color.Color
}

// Draw dessine le bouton sur l'écran
func (b *Bouton) Draw(screen *ebiten.Image) {
	// Dessiner un rectangle
	rect := ebiten.NewImage(b.Width, b.Height)
	rect.Fill(b.Color)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.X), float64(b.Y))
	screen.DrawImage(rect, op)

	// Dessiner le texte centré approximativement
	text.Draw(screen, b.Label, basicfont.Face7x13, b.X+10, b.Y+b.Height-10, color.White)
}

// IsClicked vérifie si le bouton est cliqué
func (b *Bouton) IsClicked(mouseX, mouseY int) bool {
	return mouseX >= b.X && mouseX <= b.X+b.Width &&
		mouseY >= b.Y && mouseY <= b.Y+b.Height
}
func ColorRGB(r, g, b int) color.Color{
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}