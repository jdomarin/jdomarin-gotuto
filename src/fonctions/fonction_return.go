package main

import "fmt"

// Fonction qui retourne un type int
func maxNbr(a int, b int) int {
    if a > b {
        return a // retourne la variable a de type int
    }
    return b  // retourne la variable b de type int
}

func main() {
    max := maxNbr(10, 30) // stockage du retour de la fonction dans une variable
    fmt.Println(max)

    // Utilisation directe du retour de la fonction sans la stocker dans une variable
    fmt.Printf("Valeur : %d , Type : %T\n", maxNbr(50, 30), maxNbr(50, 30))
}
