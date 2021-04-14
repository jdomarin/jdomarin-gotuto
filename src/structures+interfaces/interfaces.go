package main

import "fmt"

type Forme interface { // création de l'interface Forme
    Air() float64 // signature de la méthode Air()
    Perimetre() float64 // signature de la méthode Perimetre()
}

type Rectangle struct {
    largeur float64
    longueur float64
}

/*
Pour implémenter une interface dans Go, il suffir
d'implémenter toutes les méthodes de l'interface. Ici on
implémente la méthode Air() de l'interface Forme
*/
func (r Rectangle) Air() float64 {
    return r.largeur * r.longueur
}

/*
On implémente la méthode Perimetre() de l'interface Forme
*/
func (r Rectangle) Perimetre() float64 {
    return 2 * (r.largeur + r.longueur)
}

func main() {
    var f Forme // Initialisation de l'interface Forme
    f = Rectangle{5.0, 4.0} // affectation de la structure Rectangle à l'interface Forme
    r := Rectangle{5.0, 4.0}
    fmt.Println("Type de f :", f)
    fmt.Printf("Valeur de f : %v\n", f)
    fmt.Println("Aire du rectangle r :", f.Air())
    fmt.Println("f == r ?", f == r)
}
