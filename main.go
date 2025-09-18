package main

import (
	"Projet-Red_Zyrrathion/character"
	"Projet-Red_Zyrrathion/menu"
	"fmt"
	"time"
)

func main() {

	const (
		Reset    = "\033[0m"
		Bold     = "\033[1m"
		DarkBlue = "\033[34m"
	)

	text := "Bienvenue dans Zyrrathion !"
	for _, c := range text {
		fmt.Printf("%s%c%s", Bold+DarkBlue, c, Reset)
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Println()

	texte := "Zyrrathion est un monde melant des humains , des elfes et des nains qui copperaient pour un monde futuriste."
	for _, c := range texte {
		fmt.Printf("%s%c%s", DarkBlue, c, Reset)
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Println()

	textes := "il y a 2000 ans coexistait les monstres et les elfes mais une guerre eclata et devinrent ennemis à jamais."
	for _, c := range textes {
		fmt.Printf("%s%c%s", DarkBlue, c, Reset)
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Println()

	txt := "Vous seul a travers la contréé de Zyrrathion peuvent retablir la paix entre monstres et elfes."
	for _, c := range txt {
		fmt.Printf("%s%c%s", Bold+DarkBlue, c, Reset)
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Println()

	player := character.CharacterCreation()
	character.DisplayCharacterInfo(&player)

	menu.AfficherMenu(&player)
	fmt.Print(Bold + "Fin du jeu." + Reset)
}
