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

	"github.com/Prestonwalterbassette/Projet-Red_Zyrrathion/source/character"
	ui "github.com/Prestonwalterbassette/Projet-Red_Zyrrathion/source/interface"
	"github.com/Prestonwalterbassette/Projet-Red_Zyrrathion/source/menu"
	"github.com/hajimehoshi/ebiten/v2"
)

func GetPath(rel string) string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), rel)
}

// Affiche le texte d'intro dans la console
func afficherIntro() {
	const Reset = "\033[0m"
	const Bold = "\033[1m"
	const DarkBlue = "\033[34m"

	introTexts := []string{
		"Bienvenue dans Zyrrathion !",
		"Zyrrathion est un monde melant des humains, des elfes et des nains qui cooperaient pour un monde futuriste.",
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

func main() {
	// Créer le bouton
	fontPath := GetPath("assets/polices/Fixedsys62.ttf")
	fontFace := ui.LoadFont(fontPath, 32)

	bouton := &ui.Bouton{
		X:      860,
		Y:      880,
		Width:  200,
		Height: 50,
		Label:  "Jouer",
		Color:  ui.ColorRGB(200, 0, 0),
		Font:   fontFace,
	}

	// Lancer la fenêtre avec Ebiten
	imagePath := GetPath("assets/images/Zyrrathion.png")
	f, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	imgDecoded, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	imgEbiten := ebiten.NewImageFromImage(imgDecoded)

	ui.LancerFenetre(imgEbiten, bouton)
	if err != nil {
		fmt.Println("Fenètre fermé.")
	}

	afficherIntro()
	player := character.CharacterCreation()
	menu.AfficherMenu(&player)

}
