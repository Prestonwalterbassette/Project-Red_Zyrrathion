package combat

import (
	"github.com/Prestonwalterbassette/Projet-Red_Zyrrathion/source/experience"
	"github.com/Prestonwalterbassette/Projet-Red_Zyrrathion/source/inventory"
	"github.com/Prestonwalterbassette/Projet-Red_Zyrrathion/source/model"
	"github.com/Prestonwalterbassette/Projet-Red_Zyrrathion/source/monster"
	"fmt"
)

func DisplayMonster(m monster.Monster) {
	fmt.Println("MONSTRE")
	fmt.Println("Nom: ", m.Name)
	fmt.Println("HP: ", m.CurrentHealthPoints, "/", m.MaxHealthPoints)
	fmt.Println("Attaque: ", m.Attack)
	fmt.Println("==================")
}

func GoblinPattern(g monster.Monster, player *model.Character, rounds int) {
	for turn := 1; turn <= rounds; turn++ {
		damage := g.Attack - player.Defense%2

		if turn%3 == 0 {
			damage *= 2
		}

		player.CurrentHealthPoints -= damage
		if player.CurrentHealthPoints < 0 {
			player.CurrentHealthPoints = 0
		}

		fmt.Println(g.Name, "inflige", damage, "de dégats à", player.Name, "!")
		fmt.Println(player.Name, "HP:", player.CurrentHealthPoints, "/", player.MaxHealthPoints)
		fmt.Println()

		if player.CurrentHealthPoints <= 0 {
			fmt.Println(player.Name, "est mort !")
			break
		}
	}
}

func CharacterTurn(player *model.Character, monster *monster.Monster) {
	var choice int
	for {
		fmt.Println("=== Tour de", player.Name, " ===")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Sorts")
		fmt.Print("Choissisez une action: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			damage := player.Attack - monster.Defense%2
			monster.CurrentHealthPoints -= damage
			if monster.CurrentHealthPoints < 0 {
				monster.CurrentHealthPoints = 0
			}
			fmt.Println(player.Name, "utilise une attaque basique et fait", damage, "de dégâts !")
			fmt.Println(monster.Name, "HP: ", monster.CurrentHealthPoints, "/", monster.MaxHealthPoints)

			GoblinPattern(*monster, player, 1)
			return

		case 2:
			fmt.Println("=== Inventaire ===")
			if len(player.Inventory) == 0 {
				fmt.Println("Inventaire vide !")
			} else {
				for i, item := range player.Inventory {
					fmt.Println(i+1, ".", item)
				}
				fmt.Print("Choisissez un item à utiliser (numéro) ou 0 pour revenir : ")

				var itemChoice int
				fmt.Scanln(&itemChoice)

				if itemChoice > 0 && itemChoice <= len(player.Inventory) {
					item := player.Inventory[itemChoice-1]
					fmt.Println("Vous utilisez", item, "!")
					inventory.UseItem(player, item)

					GoblinPattern(*monster, player, 1)
					return
				}
				fmt.Println("============")
			}
		case 3:
			fmt.Println("=== Sorts Disponibles ===")
			fmt.Println("Coup de poing (8 dégats + ATK%2)")
			fmt.Println("Boule de feu (18 dégats)")
			fmt.Println("Choississez un sort:")

			var Spellchoice int
			fmt.Scanln(&Spellchoice)

			var damage int
			var manaCost int

			switch Spellchoice {
			case 1:
				damage = 8 + player.Attack/2
				manaCost = 5

			case 2:
				damage = 18 + player.Attack/2
				manaCost = 15

			default:
				fmt.Println("Sort invalide, choissisez parmi ceux que vous avez.")
				continue
			}
			if player.Mana < manaCost {
				fmt.Println("Pas assez de Mana, vous ne faites rien !")
				continue
			}

			player.Mana -= manaCost
			monster.CurrentHealthPoints -= damage
			if monster.CurrentHealthPoints < 0 {
				monster.CurrentHealthPoints = 0
			}

			fmt.Println("Vous lancez un sort et infligez", damage, "de dégâts !")
			fmt.Println(monster.Name, "HP: ", monster.CurrentHealthPoints, "/", monster.MaxHealthPoints)

			GoblinPattern(*monster, player, 1)
			return
		}
	}
}

func TrainingFight(player *model.Character) {
	goblin := monster.InitGoblin()

	fmt.Println("=== Entraînement ===")
	DisplayMonster(goblin)

	Turn := 1

	for player.CurrentHealthPoints > 0 && goblin.CurrentHealthPoints > 0 {
		fmt.Println("Tour n° ", Turn, "!")

		if player.Initiative >= goblin.Initiative {
			CharacterTurn(player, &goblin)
			if goblin.CurrentHealthPoints <= 0 {
				fmt.Println(goblin.Name, "A été vaincu !")

				experience.GainExp(player, goblin)
				break
			}
		} else {
			GoblinPattern(goblin, player, 1)
			if player.CurrentHealthPoints <= 0 {
				fmt.Print(player.Name, "est mort...")
				break
			}
			CharacterTurn(player, &goblin)
			if goblin.CurrentHealthPoints <= 0 {
				fmt.Print(goblin.Name, "Est vaincu !")

				experience.GainExp(player, goblin)
				break
			}
		}

		Turn++
	}
	fmt.Println("Fiiin du combat !")
	if player.CurrentHealthPoints <= 0 {
		fmt.Println("Vous etes mort...")
	} else {
		fmt.Println("Vous survivez à l'entraînement ! Bravo !")
		fmt.Println("Votre Expérience actuelle:", player.Experience, "/", player.MaxExperience)
	}
}
