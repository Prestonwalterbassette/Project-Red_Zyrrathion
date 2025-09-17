package equipment

import (
	"Projet-Red_Zyrrathion/character"
	"fmt"
)

// Structure représentant un équipement
type Equipment struct {
	Name    string
	Slot    string // Head, Torso, Feet
	BonusHP int
}

// Initialise les équipements de base du joueur
func InitEquipment(c *character.Character) {
	// Équipements de base
	baseEquipment := map[string]Equipment{
		"Head":  {"Casque en cuir", "Head", 15},
		"Torso": {"Plastron en fer", "Torso", 20},
		"Feet":  {"Bottes solides", "Feet", 10},
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
func DisplayEquipment(c *character.Character) {
	if len(c.EquipmentSlots) == 0 {
		fmt.Println("Equipements: Aucun !")
		return
	}

	fmt.Println("Équipements équipés :")
	for slot, eq := range c.EquipmentSlots {
		fmt.Printf(" - %s : %s (+%d PV)\n", slot, eq.Name, eq.BonusHP)
	}
}
