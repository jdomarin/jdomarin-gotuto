package main

import (
    "fmt"
)

func verificationDivision(nbr float64) (float64, bool) {
    if nbr <= 0 {
        return 0, false
    } else {
        return nbr,true
    }
}

func main() {
    nbr := 0.0

    nbr, err := verificationDivision(nbr)

    if err == false {
        panic("Erreur : Il est impossible de diviser par 0 !")
    } else {
        fmt.Println("Aucune erreur")
    }
}
