package inventory

import (
	"fmt"
	"github.com/Prestonwalterbassette/Projet-Red_Zyrrathion/source/model"
)

func AccessInventory(c *model.Character) {
	if len(c.Inventory) == 0 {
		fmt.Println("L'inventaire est vide.")
		return
	}

	fmt.Println("==== Inventaire de", c.Name, " ====")
	for i, item := range c.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	fmt.Printf("===================")
}

func takePot(potion int, hp *int, maxHP int) {
	*hp += potion
	if *hp > maxHP {
		*hp = maxHP
	}
}

func UseItem(c *model.Character, item string) {
	switch item {
	case "Livre de sort : Boule de Feu":
		fmt.Println("Vous utilisez le Livre de sort : Boule de Feu !")
		fmt.Println(c.Name, "apprend le sort : Boule de Feu !")

		for i, it := range c.Inventory {
			if it == item {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
				break
			}
		}
	case "Potion de vie":
		fmt.Println("Vous utilisez une Potion de vie.")
		heal := 50
		takePot(heal, &c.CurrentHealthPoints, c.MaxHealthPoints)
		fmt.Println(c.Name, "a maintenant", c.CurrentHealthPoints, "/", c.MaxHealthPoints, "HP.")

		for i, it := range c.Inventory {
			if it == item {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
				break
			}
		}
	default:
		fmt.Println("Vous utilisez:", item)
	}
}

func AddItem(c *model.Character, item string) {
	if len(c.Inventory) >= 10 {
		fmt.Println("Inventaire plein ! Vous ne pouvez portez plus de 10 objets dans votre inventaire !")
		return
	}
	c.Inventory = append(c.Inventory, item)
	fmt.Println(item, "a été ajouté à l'inventaire !")
}
