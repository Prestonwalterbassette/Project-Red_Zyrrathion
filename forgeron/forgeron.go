package forgeron

import "fmt"

// Item représente un objet fabriquable
type Item struct {
	Nom                string
	RessourcesRequises map[string]int
}

// Personne représente un joueur avec inventaire et items fabriqués
type Personne struct {
	Money      int
	Inventory map[string]int
	Items      []string
}

// Fabriquer tente de fabriquer un objet à partir de l'inventaire
func (p *Personne) Fabriquer(item Item) {
	for res, qte := range item.RessourcesRequises {
		if p.Inventory[res] < qte {
			fmt.Printf("Pas assez de %s pour fabriquer %s\n", res, item.Nom)
			return
		}
	}
	for res, qte := range item.RessourcesRequises {
		p.Inventory[res] -= qte
	}
	p.Items = append(p.Items, item.Nom)
	fmt.Println("Item fabriqué :", item.Nom)
}

// ObjetsAFabriquer contient tous les objets disponibles pour fabrication
var ObjetsAFabriquer = map[string]Item{
	"Chapeau d'aventurier": {
		Nom: "Chapeau d'aventurier",
		RessourcesRequises: map[string]int{
			"plume de corbeau": 1,
			"cuir de sanglier": 1,
		},
	},
	"Tunique de l'aventurier": {
		Nom: "Tunique de l'aventurier",
		RessourcesRequises: map[string]int{
			"fourrure de loup": 2,
			"peau de troll":    1,
		},
	},
	"Bottes de l'aventurier": {
		Nom: "Bottes de l'aventurier",
		RessourcesRequises: map[string]int{
			"fourrure de loup": 1,
			"cuir de sanglier": 1,
		},
	},
}

// AfficherObjetsAFabriquer affiche tous les objets disponibles et leurs ressources
func AfficherObjetsAFabriquer() {
	fmt.Println("Objets disponibles à fabriquer")
	for nom, item := range ObjetsAFabriquer {
		fmt.Println("- ", nom, "requiert ")
		for res, qte := range item.RessourcesRequises {
			fmt.Printf("    %d x %s\n", qte, res)
		}
	}
}
