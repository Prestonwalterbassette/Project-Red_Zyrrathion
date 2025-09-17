package main

import(
	"fmt"
	"Projet-Red_Zyrrathion/character"
	"Projet-Red_Zyrrathion/menu"
)

func main() {
	fmt.Println("Bienvenu dans Zyrrathion !")
	player := character.CharacterCreation()
	character.DisplayCharacterInfo(&player)

	menu.AfficherMenu(&player)
	fmt.Print("Fin du jeu.")
}