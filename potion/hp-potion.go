package main

import "fmt"

func takePot(potion int, hp *int, maxHp int) {
	*hp += potion
	if *hp > maxHp {
		*hp = maxHp
	}
}

func main() {
	playerHp := 40
	maxHp := 80
	hpPotion := 50
	takePot(hpPotion, &playerHp, maxHp)
	fmt.Println("Your HP after taking the potion is:", playerHp)
}
