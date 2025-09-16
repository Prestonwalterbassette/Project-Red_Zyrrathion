package character

import "fmt"

type Character struct {
	Name                string   `json:"name"`
	Gender              string   `json:"gender"`
	Race                string   `json:"race"`
	Class               string   `json:"class"`
	Level               int      `json:"level"`
	Experience          int      `json:"EXP"`
	MaxExperience       int      `json:"MaxEXP"`
	Skills              []string `json:"skills"`
	CurrentHealthPoints int      `json:"HP"`
	MaxHealthPoints     int      `json:"maxHP"`
	Attack              int      `json:"ATK"`
	Defense             int      `json:"DEF"`
	Mana                int      `json:"MP"`
	MaxMana             int      `json:"MaxMP"`
	Inventory           []string `json:"inventory"`
	Resources   		map[string]int `json:"resources"`
	Items 				[]string `json:"items"`
	Initiative          int      `json:"initiative"`
	Currency 			int 	 `json:"Money"`
	Purchasehistory 	[]int 	 `json:"History"`
	EquipmentSlots 		map[string]string `json:"Equipment"`
}

func InitCharacter(name, gender, race, class string, level int, experience int, maxexperience int, skills []string, hp, maxHP int, ATK int, DEF int, mana int, maxmana int, inventory []string, initiative int, currency int) Character {
	hasBaseSkill := false
	for _, s := range skills {
		if s == "Coup de poing" {
			hasBaseSkill = true
			break
		}
	}
	if !hasBaseSkill {
		skills = append(skills, "Coup de poing")
	}

	return Character{
		Name:                name,
		Gender:              gender,
		Race:                race,
		Class:               class,
		Level:               1,
		Experience:          0,
		Skills:              skills,
		CurrentHealthPoints: 40,
		MaxHealthPoints:     80,
		Attack:              5,
		Defense:             5,
		Mana:                25,
		MaxMana:             25,
		Inventory:           inventory,
		Initiative:          10,
		Currency:            100,
	}
}

func CharacterCreation() Character {
	var name string
	var maxHP int
	var race string
	var ATK int
	var DEF int
	var class string
	var mana int
	var maxmana int

	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scanln(&name)

	if len(name) > 0 {
		first := string(name[0])
		if name[0] >= 'a' && name[0] <= 'z' {
			first = string(name[0] - 32)
		}
		if len(name) > 1 {
			name = first + name[1:]
		} else {
			name = first
		}
	}

	for {
		fmt.Print("Choissisez une race (Humain, Elfe, Nain) : ")
		fmt.Scanln(&race)

		switch race {
		case "Humain", "humain", "HUMAIN":
			race = "Humain"
			maxHP = 100
			ATK = 5
			DEF = 10
			maxmana = 25

		case "Elfe", "elfe", "ELFE":
			race = "Elfe"
			maxHP = 80
			ATK = 4
			DEF = 9
			maxmana = 50

		case "Nain", "nain", "NAIN":
			race = "Nain"
			maxHP = 120
			ATK = 6
			DEF = 11
			maxmana = 15

		default:
			fmt.Println("Race invalide, veuillez réessayer.")
			continue
		}
		break
	}
	for {
		fmt.Println("Choissisez une class(Guerrier, Archer, Mage) pour votre", race, " .")
		fmt.Scanln(&class)

		switch class {
		case "Guerrier", "GUERRIER", "guerrier":
			class = "Guerrier"
			text := "En levant ton épée, tu montres au monde que tu es une personne dignes d'éloges et de respect."
			fmt.Print(text)
			maxHP = maxHP + 50
			ATK = ATK + 5
			DEF = DEF + 6
			maxmana += 15

		case "Archer", "ARCHER", "archer":
			class = "Archer"
			text := "Ton arc montre aux civilisations tes compétences et ta précision, tu ne rates aucune cible de ton regard."
			fmt.Print(text)
			maxHP = maxHP + 25
			ATK = ATK + 7
			DEF = DEF + 3
			maxmana += 20

		case "Mage", "MAGE", "mage":
			class = "Mage"
			text := "Tes sorts et ta mana sont unique et inégalé. Tu sème le chaos et la destruction avec tes sorts."
			fmt.Print(text)
			maxHP = maxHP + 15
			ATK = ATK + 10
			DEF = DEF + 2
			maxmana += 50

		default:
			fmt.Print("Invalid. Choissisez votre class parmi ceux montrés.")
			continue
		}
		break
	}

	hp := maxHP / 2
	mana = maxmana / 2

	player := Character{
		Name:          name,
		Gender:        "Male",
		Race:          race,
		Class:         class,
		Level:         1,
		Experience:    0,
		MaxExperience: 100,

		Skills:              []string{"Coup de poing"},
		CurrentHealthPoints: hp,
		MaxHealthPoints:     maxHP,
		Attack:              ATK,
		Defense:             DEF,
		Mana:                mana,
		MaxMana:             maxmana,
		Inventory:           []string{"Potion"},
		Initiative:          10,
		Currency:            100,
	}

	fmt.Println("Personnage crée :", player.Name, " - ", player.Race)
	return player
}

func DisplayCharacterInfo(c Character) {

	fmt.Println("Character Information:")
	fmt.Println("Name:", c.Name)
	fmt.Println("Gender:", c.Gender)
	fmt.Println("Race:", c.Race)
	fmt.Println("Class:", c.Class)
	fmt.Println("Level:", c.Level)
	fmt.Println("EXP:", c.Experience, "/", c.MaxExperience)
	fmt.Println("Health Points:", c.CurrentHealthPoints, "/", c.MaxHealthPoints)
	fmt.Println("ATK:", c.Attack)
	fmt.Println("DEF:", c.Defense)
	fmt.Println("MP:", c.Mana, "/", c.MaxMana)
	fmt.Println("Initiative:", c.Initiative)
	fmt.Println("-------------------------")

	if len(c.Skills) > 0 {
		fmt.Println("Skills:", c.Skills)
	} else {
		fmt.Println("Skills: None")
	}

	if len(c.Inventory) > 0 {
		fmt.Println("Inventory:", c.Inventory)
	} else {
		fmt.Println("Inventory: Empty")
	}
	fmt.Println("-------------------------")
}

func (c *Character) Spellbook(spell string) {
	for _, s := range c.Skills {
		if s == spell {
			fmt.Println("Le sort", spell, "est déjà appris !")
			return
		}
	}
	c.Skills = append(c.Skills, spell)
	fmt.Println("Le sort", spell, "a été ajouté à votre liste de sorts !")
}
