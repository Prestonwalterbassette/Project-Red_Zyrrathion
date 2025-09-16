package menu

import (
	"fmt"
	"Projet-Red_Zyrrathion/character"
	"Projet-Red_Zyrrathion/combat"
)

func AfficherMenu(c *character.Character) {
	var choice int

	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Entraînement")
		fmt.Println("4. Quitter")
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			character.DisplayCharacterInfo(*c)
		case 2:
			fmt.Println("Inventaire :", c.Inventory)
		case 3:
			combat.TrainingFight(c)
		case 4:
			fmt.Println("Voulez-vous quitter le jeu aussi tôt ? (Y/N)")
			var input string
			fmt.Scanln(&input)
			if input == "Y" || input == "y" {
				fmt.Println("Adieu !")
				return
			}
		default:
			fmt.Println("Choix invalide !")
		}
	}
}
