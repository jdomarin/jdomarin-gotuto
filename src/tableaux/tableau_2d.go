package main

import "fmt"

func main() {
    const(
        maxLigne   int = 3
        maxColonne int = 3
    )

    var doubleTableau [maxLigne][maxColonne]int

    fmt.Println(doubleTableau)

    fmt.Println("-------------------")

    // modification de la ligne 3, colonne 2
    doubleTableau[2][1] = 5 // modification du 2ème élément du 3ème tableau
    fmt.Println(doubleTableau)

    chaine := "Hello"
    fmt.Println(chaine)
}
