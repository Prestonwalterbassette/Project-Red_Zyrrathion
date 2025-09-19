package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"Projet-Red_Zyrrathion/character"
	ui "Projet-Red_Zyrrathion/interface"
	"Projet-Red_Zyrrathion/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	background *ebiten.Image
	mouseDown  bool
	started    bool
	bouton     *ui.Bouton
}

func GetPath(rel string) string {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)
	return filepath.Join(baseDir, rel)
}

// Affiche le texte d'intro dans la console
func afficherIntro() {
	const (
		Reset    = "\033[0m"
		Bold     = "\033[1m"
		DarkBlue = "\033[34m"
	)

	introTexts := []string{
		"Bienvenue dans Zyrrathion !",
		"Zyrrathion est un monde melant des humains, des elfes et des nains qui copperaient pour un monde futuriste.",
		"Il y a 2000 ans coexistait les monstres et les elfes mais une guerre eclata et devinrent ennemis à jamais.",
		"Vous seul à travers la contrée de Zyrrathion pouvez retablir la paix entre monstres et elfes.",
	}

	for i, txt := range introTexts {
		for _, c := range txt {
			if i == 0 || i == 3 {
				fmt.Printf("%s%s%c%s", Bold, "\033[34m", c, Reset)
			} else {
				fmt.Printf("%s%c%s", DarkBlue, c, Reset)
			}
			time.Sleep(30 * time.Millisecond)
		}
		fmt.Println()
	}
}

func (g *Game) Update() error {
	if g.started {
		// Le jeu a commencé, aucune action ici
		return nil
	}

	x, y := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !g.mouseDown {
			g.mouseDown = true
			if g.bouton.IsClicked(x, y) {
				fmt.Println("Bouton 'Jouer' cliqué !")
				g.started = true

				// Affichage du texte d'intro
				afficherIntro()

				// Création du personnage et menu
				player := character.CharacterCreation()
				menu.AfficherMenu(&player)
				fmt.Println("Fin du jeu.")
			}
		}
	} else {
		g.mouseDown = false
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Dessiner l'image de fond
	op := &ebiten.DrawImageOptions{}
	sw, sh := g.background.Size()
	scaleX := float64(1920) / float64(sw)
	scaleY := float64(1080) / float64(sh)
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(g.background, op)

	// Dessiner le bouton si le jeu n'a pas commencé
	if !g.started {
		g.bouton.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1920, 1080
}

func main() {
	// Charger l'image de fond
	imagePath := GetPath("../assets/images/Zyrrathion.png")
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	fontPath := GetPath("../assets/polices/Fixedsys62.ttf")
	fontFace := ui.LoadFont(fontPath, 32)

	// Initialiser le jeu et le bouton
	boutonJouer := &ui.Bouton{
		X:      860,
		Y:      880,
		Width:  200,
		Height: 50,
		Label:  "Jouer",
		Color:  ui.ColorRGB(200, 0, 0),
		Font:   fontFace,
	}

	game := &Game{
		background: ebiten.NewImageFromImage(img),
		mouseDown:  false,
		started:    false,
		bouton:     boutonJouer,
	}

	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Zyrrathion")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
