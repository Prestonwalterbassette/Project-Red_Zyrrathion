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
	Font 		  font.Face
}

func LoadFont(path string, size float64) font.Face{
	data, err := ioutil.ReadFile(path)
	if err != nil{
		log.Fatal(err)
	}

	tt, err := opentype.Parse(data)
	if err != nil{
		log.Fatal(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 	size,
		DPI: 	72,
		Hinting: font.HintingFull,
	})
	if err != nil{
		log.Fatal(err)
	}
	
	return face
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
	text.Draw(screen, b.Label, b.Font, b.X+10, b.Y+b.Height-10, color.White)
}

// IsClicked vérifie si le bouton est cliqué
func (b *Bouton) IsClicked(mouseX, mouseY int) bool {
	return mouseX >= b.X && mouseX <= b.X+b.Width &&
		mouseY >= b.Y && mouseY <= b.Y+b.Height
}
func ColorRGB(r, g, b int) color.Color{
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}