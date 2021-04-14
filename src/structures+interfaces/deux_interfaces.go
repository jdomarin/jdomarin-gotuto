package main

import (
    "fmt"
    "math"
)

type Forme interface { // création de l'interface Forme
    Air() float64 // signature de la méthode Air()
    Perimetre() float64 // signature de la méthode Perimetre()
}

/*----------Début Structure Rectangle----------*/
type Rectangle struct {
    largeur float64
    longueur float64
}

func (r Rectangle) Air() float64 {
    return r.largeur * r.longueur
}

func (r Rectangle) Perimetre() float64 {
    return 2 * (r.largeur + r.longueur)
}
/*----------Fin Structure Rectangle----------*/


/*----------Début Structure Cercle----------*/
type Cercle struct {
    rayon float64
}

func (c Cercle) Air() float64 {
    return math.Pi * c.rayon * c.rayon
}

func (c Cercle) Perimetre() float64 {
    return 2 * math.Pi * c.rayon
}
/*----------Fin Structure Cercle----------*/

// Fonction qui prend en compte un paramètre de type Cercle ou Rectangle
func AirPerimetrePresentation(f Forme) {
    fmt.Println("- Aire:", f.Air())
    fmt.Println("- Périmètre :", f.Perimetre())
}


func main() {
    var r Forme = Rectangle{5.0, 4.0}
    var c Forme = Cercle{5.0}

    fmt.Println("Cercle :")
    AirPerimetrePresentation(c)
    fmt.Println("\nrectangle :")
    AirPerimetrePresentation(r)
}
