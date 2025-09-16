package main

import (
	"fmt"
	"time"
)

func poisonPot(hp *int) {
	for i := 0; i < 3; i++ {
		*hp -= 10
		if *hp < 0 {
			*hp = 0
		}
		fmt.Println("second %d: You HP is now %d\n", i+1, *hp)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	playerHp := 80
	fmt.Println("starting HP:", playerHp)
	poisonPot(&playerHp)
	fmt.Println("Your HP after taking the poison potion is:", playerHp)
}
