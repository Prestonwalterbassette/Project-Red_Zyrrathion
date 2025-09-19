package equipment

import (
	"github.com/Prestonwalterbassette/Projet-Red_Zyrrathion/source/model"
	"fmt"
)

// Initialise les équipements de base du joueur
func InitEquipment(c *model.Character) {
	// Équipements de base
	baseEquipment := map[string]model.Equipment{
		"Head": {
			Name:  "Casque en cuir",
			Slot: "Head",
			BonusHP: 15,
	},
		"Torso": {
			Name: "Plastron en fer",
			Slot: "Torso",
			BonusHP: 20,
		},
		"Feet": {
			Name: "Bottes solides",
			Slot: "Feet",
			BonusHP: 10,
		},
	}

	// Application des bonus PV
	bonusHP := map[string]int{
		"Head":  15,
		"Torso": 20,
		"Feet":  10,
	}

	c.EquipmentSlots = baseEquipment

	for slot, eq := range baseEquipment {
		c.MaxHealthPoints += eq.BonusHP
		c.CurrentHealthPoints += eq.BonusHP
		fmt.Print(slot, "équipé:", eq.Name, "+", bonusHP[slot], "de HP.")
	}
}

// Affichage des équipements du joueur
func DisplayEquipment(c *model.Character) {
	if len(c.EquipmentSlots) == 0 {
		fmt.Println("Equipements: Aucun !")
		return
	}

	fmt.Println("Équipements équipés :")
	for slot, eq := range c.EquipmentSlots {
		fmt.Printf(" - %s : %s (+%d PV)\n", slot, eq.Name, eq.BonusHP)
	}
}
