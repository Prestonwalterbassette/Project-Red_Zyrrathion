package status

import "fmt"

type Character struct {
	Name  string
	HP    int
	MaxHP int
}

func IsDead(c *Character) bool {
	if c.HP <= 0 {
		fmt.Println("You are Dead!")
		c.HP = c.MaxHP / 2
		fmt.Println("You have been resurrected with", c.HP, "HP!")
		return true
	} else {
		fmt.Println("You are Alive with", c.HP, "HP")
		return false
	}
}