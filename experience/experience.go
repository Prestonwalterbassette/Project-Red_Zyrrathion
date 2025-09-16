package experience

import (
	"fmt"
	"Projet-Red_Zyrrathion/character"
	"Projet-Red_Zyrrathion/monster"
)

func LevelUp(player *character.Character) {
	player.Level++
	fmt.Println("Vous montez d'un niveau ! Bravo ! Vous ètes niveau:", player.Level)

	player.Attack += 2
	player.Defense += 1
	player.CurrentHealthPoints += 5 

	player.MaxExperience += 50
}
func GainExp(player *character.Character, monster monster.Monster) {
	fmt.Println("Vous gagnez:", monster.ExpReward, "EXP !")
	player.Experience += monster.ExpReward

	for player.Experience >= player.MaxExperience {
		player.Experience -= player.MaxExperience
		player.Level++
	}
}