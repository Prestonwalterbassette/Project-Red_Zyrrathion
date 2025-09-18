package model

type Character struct {
	Name                string               `json:"name"`
	Gender              string               `json:"gender"`
	Race                string               `json:"race"`
	Class               string               `json:"class"`
	Level               int                  `json:"level"`
	Experience          int                  `json:"EXP"`
	MaxExperience       int                  `json:"MaxEXP"`
	Skills              []string             `json:"skills"`
	CurrentHealthPoints int                  `json:"HP"`
	MaxHealthPoints     int                  `json:"maxHP"`
	Attack              int                  `json:"ATK"`
	Defense             int                  `json:"DEF"`
	Mana                int                  `json:"MP"`
	MaxMana             int                  `json:"MaxMP"`
	Inventory           []string             `json:"inventory"`
	Resources           map[string]int       `json:"resources"`
	Items               []string             `json:"items"`
	Initiative          int                  `json:"initiative"`
	Currency            int                  `json:"Money"`
	Purchasehistory     []int                `json:"History"`
	EquipmentSlots      map[string]Equipment `json:"Equipment"`
}

type Equipment struct {
	Name    string
	Slot    string // Head, Torso, Feet
	BonusHP int
}

type Personne struct {
	Money     int
	Inventory map[string]int
	Items     []string
}

type Monster struct {
	Name                string `json:"name"`
	MaxHealthPoints     int    `json:"MaxHP"`
	CurrentHealthPoints int    `json:"HP"`
	Attack              int    `json:"ATK"`
	Defense             int    `json:"DEF"`
	ExpReward           int    `json:"EXPreward"`
	Initiative          int    `json:"initiative"`
}

type Argent struct {
	Or int
}
