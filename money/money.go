package main

import "fmt"

// Définition d'une personne
type Personne struct {
	nom        string // anciennement lunareth
	argent     int    // argent de la personne
	historique []int  // historique des achats
	personnage string // attribut supplémentaire
}

// Définition d'une structure Argent
type Argent struct {
	or int
}

func main() {
	// Initialisation
	p := Personne{nom: "Lunareth", personnage: "Elfe", argent: 100}
	m := Argent{or: 100}

	fmt.Println(p)
	fmt.Println(m)
	fmt.Println("Argent initial :", p.argent, m.or)

	// Exemple d’achats
	acheter(&p, 3)
	fmt.Println("Argent après potion :", p.argent)
	fmt.Println("Historique des achats :", p.historique)

	acheter(&p, 6)
	fmt.Println("Argent après potion de poison :", p.argent)
	fmt.Println("Historique des achats :", p.historique)

	acheter(&p, 25)
	fmt.Println("Argent après livre magique :", p.argent)
	fmt.Println("Historique des achats :", p.historique)

	acheter(&p, 7)
	fmt.Println("Argent après peau de troll :", p.argent)
	fmt.Println("Historique des achats :", p.historique)
}

// Fonction pour effectuer un achat
func acheter(p *Personne, montant int) {
	if p.argent >= montant {
		p.argent -= montant
		p.historique = append(p.historique, montant)
		fmt.Println("Achat effectué :", montant, "pièces d or")
	} else {
		fmt.Println("Fonds insuffisants pour cet achat :", montant, "pièces d or")
	}
}
