package main

import "fmt"

// création de notre structure Personnage
type Personnage struct {
    nom        string
    vie        int
    puissance  int
    mort       bool
    inventaire [3]string
}

func main() {
    // initialisation de ma structure Personnage
    var p1 Personnage
    var p2 Personnage

    p1.Init("barbare", 200, 20, false, [3]string{"épée", "bouclier", "armure"})
    p2.Init("magicien", 100, 40, false, [3]string{"potions", "poisons", "bâton"})

    p1.affichage()
    p2.affichage()
}

/*
@description : Affiche des informations sur un personnage

@return : void
*/
func (p Personnage) affichage() {
    fmt.Println("----------------------------------------------------")
    fmt.Println("Vie du personnage", p.nom, ":", p.vie)
    fmt.Println("Puissance du personnage", p.nom, ":", p.puissance)

    if p.mort {
        fmt.Println("Vie du personnage", p.nom, "est mort")
    } else {
        fmt.Println("Vie du personnage", p.nom, "est vivant")
    }

    fmt.Println("\nLe personnage", p.nom, "possède dans son inventaire :", p.vie)

    for _, item := range p.inventaire {
        fmt.Println("-", item)
    }
}


/*
@description : Surcharge des valeurs par défaut

@return: void
*/
func (p *Personnage) Init(nom string, vie int, puissance int, mort bool, inventaire [3]string) {
    p.nom = nom
    p.vie = vie
    p.puissance = puissance
    p.mort = mort
    p.inventaire = inventaire
}
