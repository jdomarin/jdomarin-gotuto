package main

import "fmt"

func main() {
    a := 5
    b := 8
    fmt.Println("Avant fonction a =", a, " b =", b)
    a, b = additionTrois(a, b) // stockage du retour de la fonction dans deux variables
    fmt.Println("Après fonction a =", a, " b =", b)
}

// retourne deux types int
func additionTrois(a int, b int) (int, int) {
    return a + 3, b + 3
}
