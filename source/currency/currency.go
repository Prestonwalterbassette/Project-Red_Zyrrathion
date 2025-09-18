package currency

import  "fmt"

// Définition d'une personne
type Personne struct {
	Nom        string // anciennement lunareth
	Argent     int    // argent de la personne
	Historique []int  // historique des achats
	Personnage string // attribut supplémentaire
}

// Définition d'une structure Argent
type Argent struct {
	or int
}

// Fonction pour effectuer un achat
func Acheter(p *Personne, montant int) {
	if p.Argent >= montant {
		p.Argent -= montant
		p.Historique = append(p.Historique, montant)
		fmt.Println("Achat effectué :", montant, "pièces d or")
	} else {
		fmt.Println("Fonds insuffisants pour cet achat :", montant, "pièces d or")
	}
}
