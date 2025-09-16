package monster

type Monster struct{
	Name                string `json:"name"`
	MaxHealthPoints     int    `json:"MaxHP"`
	CurrentHealthPoints int    `json:"HP"`
	Attack              int    `json:"ATK"`
	Defense				int	   `json:"DEF"`
	ExpReward			int	   `json:"EXPreward"`
	Initiative          int    `json:"initiative"`
}

func InitGoblin() Monster {
	return Monster {
		Name:                "Goblin d'entraînement",
		MaxHealthPoints:     40,
		CurrentHealthPoints: 40,
		Attack:              5,
		Defense:			 3,
		ExpReward:           15,
		Initiative:          6,
	}
}