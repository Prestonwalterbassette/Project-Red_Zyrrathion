package menu

import (
	"Projet-Red_Zyrrathion/character"
	"Projet-Red_Zyrrathion/combat"
	"Projet-Red_Zyrrathion/equipment"
	"Projet-Red_Zyrrathion/forgeron"
	"Projet-Red_Zyrrathion/merchant"
	"Projet-Red_Zyrrathion/model"
	"fmt"
	"os"
)

func AfficherMenu(c *model.Character) {
	var choice int

	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Entraînement")
		fmt.Println("4. Forgeron")
		fmt.Println("5. Marchant")
		fmt.Println("6. Equipement")
		fmt.Println("7. Quitter")
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			character.DisplayCharacterInfo(c)
		case 2:
			fmt.Println("Inventaire :", c.Inventory)
		case 3:
			combat.TrainingFight(c)
		case 4:
			fmt.Println("FORGERON")
			forgeron.AfficherObjetsAFabriquer()

			fmt.Println("Quel objet veux-tu fabriquer ? Soit rapide.")
			fmt.Println("Tape le nom exact. Ou 0 pour revenir.")
			var objet string
			fmt.Scanln(&objet)

			if objet == "0" {
				break
			}

			item, ok := forgeron.ObjetsAFabriquer[objet]
			if !ok {
				fmt.Println("De quoi tu me parles ?")
				break
			}

			// on crée un forgeron.Personne temporaire
			personne := forgeron.Personne{
				Money:     c.Currency,
				Inventory: c.Resources, // <-- le map[string]int pour les ressources
				Items:     c.Items,
			}

			personne.Fabriquer(item)

			// on remet à jour le character
			c.Currency = personne.Money
			c.Resources = personne.Inventory
			c.Items = personne.Items

		case 5:
			fmt.Println("MARCHANT")
			merchant.Marchand(c)

		case 6:
			fmt.Println("Equipement")
			equipment.DisplayEquipment(c)

		case 7:
			fmt.Println("Voulez-vous quitter le jeu aussi tôt ? (Y/N)")
			var input string
			fmt.Scanln(&input)
			if input == "Y" || input == "y" {
				fmt.Println("Adieu !")
				os.Exit(0)
			}
		default:
			fmt.Println("Choix invalide !")
		}
	}
}
