package main 

import "fmt"

// Structure représentant un équipement
type Equipment struct {
	Name    string
	Slot    string // Head, Torso, Feet
	BonusHP int
}

// Initialise les équipements de base du joueur
func InitEquipment(c *attributescharacter) {
	// Équipements de base
	equipment := map[string]Equipment{
		"Head":  {"Casque en cuir", "Head", 15},
		"Torso": {"Plastron en fer", "Torso", 20},
		"Feet":  {"Bottes solides", "Feet", 10},
	}

	// Application des bonus PV
	for _, eq := range equipment {
		c.MaxHealthPoints += eq.BonusHP
		c.CurrentHealthPoints += eq.BonusHP
	}

	c.EquipmentSlots = equipment
}

// Affichage des équipements du joueur
func DisplayEquipment(c attributescharacter)
		fmt.Println("Équipements équipés :")
		for slot, eq := range c.EquipmentSlots {
			fmt.Printf(" - %s : %s (+%d PV)\n", slot, eq.Name, eq.BonusHP)
		}
	 else {
		fmt.Println("Équipements : Aucun")
	}
	