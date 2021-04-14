package personnage

import "fmt"

type Personnage struct {
    Nom        string
    Vie        int
    Puissance  int
    Mort       bool
    Inventaire [3]string
}

/*
@description : Affiche des informations sur un personnage

@return : void
*/
func (p Personnage) Affichage() { // déclaration de ma méthode Affichage() liée à ma structure Personnage
    fmt.Println("----------------------------------------------------")
    fmt.Println("Vie du personnage", p.Nom, ":", p.Vie)
    fmt.Println("Puissance du personnage", p.Nom, ":", p.Puissance)

    if p.Mort {
        fmt.Println("Vie du personnage", p.Nom, "est mort")
    } else {
        fmt.Println("Vie du personnage", p.Nom, "est vivant")
    }

    fmt.Println("\nLe personnage", p.Nom, "possède dans son inventaire :", p.Vie)

    for _, item := range p.Inventaire {
        fmt.Println("-", item)
    }
}


/*
@description : Créer une instance de la classe Personnage

@return: struct Personnage
*/
func New(Nom string, Vie int, Puissance int, Mort bool, Inventaire [3]string) Personnage {
    personnage := Personnage{Nom, Vie, Puissance, Mort, Inventaire}
    return personnage
}
