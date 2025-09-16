package main

import (
	"fmt"

	"Projet-Red_Zyrrathion/character"
	"Projet-Red_Zyrrathion/menu"
)

func main() {
	fmt.Println("Bienvenue dans Zyrrathion !")

	// Création du personnage
	player := character.CharacterCreation()

	// Initialise les champs nécessaires
	player.Currency = 100
	player.Purchasehistory = []int{}
	player.Inventory = []string{}
	player.Resources = make(map[string]int)
	player.Items = []string{}

	// Affiche les infos du personnage
	character.DisplayCharacterInfo(player)

	// Lancer le menu principal
	menu.AfficherMenu(&player)

	fmt.Println("Fin du jeu.")
}
