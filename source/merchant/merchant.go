package merchant

import (
	"fmt"
	"github.com/Prestonwalterbassette/Projet-Red_Zyrrathion/source/model"
	"github.com/Prestonwalterbassette/Projet-Red_Zyrrathion/source/currency"
)

// Marchand lance le menu marchand pour acheter des objets
func Marchand(c *model.Character) {
	for {
		fmt.Println("\n=== MARCHAND ===")
		fmt.Println("1. Acheter Potion (3 pièces)")
		fmt.Println("2. Acheter Potion de poison (6 pièces)")
		fmt.Println("3. Acheter Livre magique (25 pièces)")
		fmt.Println("4. Acheter Peau de troll (7 pièces)")
		fmt.Println("0. Retour au menu")
		fmt.Printf("Argent disponible : %d pièces\n", c.Currency)
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scanln(&choix)

		// Crée une personne temporaire pour utiliser le package currency
		person := &currency.Personne{
			Nom:        c.Name,
			Personnage: c.Race,
			Argent:     c.Currency,
			Historique: c.Purchasehistory, // si tu as déjà un historique
		}

		acheter := func(montant int, nom string) {
			currency.Acheter(person, montant) // utilise ton package currency
			// si l'achat a été fait, ajoute l'objet dans l'inventaire
			if c.Currency != person.Argent { 
				c.Inventory = append(c.Inventory, nom)
				c.Currency = person.Argent // synchronise avec Character
				c.Purchasehistory = person.Historique
				fmt.Printf("%s ajouté à l'inventaire.\n", nom)
			}
		}

		switch choix {
		case 0:
			return // quitte le marchand
		case 1:
			acheter(3, "Potion")
		case 2:
			acheter(6, "Potion de poison")
		case 3:
			acheter(25, "Livre magique")
		case 4:
			acheter(7, "Peau de troll")
		default:
			fmt.Println("Choix invalide !")
		}
	}
}
