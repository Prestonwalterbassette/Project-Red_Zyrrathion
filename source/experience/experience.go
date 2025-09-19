package experience

import (
	"fmt"
	"github.com/Prestonwalterbassette/Projet-Red_Zyrrathion/source/monster"
	"github.com/Prestonwalterbassette/Projet-Red_Zyrrathion/source/model"
)

func LevelUp(player *model.Character) {
	player.Level++
	fmt.Println("Vous montez d'un niveau ! Bravo ! Vous ètes niveau:", player.Level)

	player.Attack += 2
	player.Defense += 1
	player.CurrentHealthPoints += 5 

	player.MaxExperience += 50
}
func GainExp(player *model.Character, monster monster.Monster) {
	fmt.Println("Vous gagnez:", monster.ExpReward, "EXP !")
	player.Experience += monster.ExpReward

	for player.Experience >= player.MaxExperience {
		player.Experience -= player.MaxExperience
		player.Level++
	}
}