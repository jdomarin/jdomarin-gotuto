package main

import "fmt"

func main() {
    mois := []string{"Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet"}

    fmt.Println(mois)

    // index à supprimer de notre slice
    indexASupprimer := 1

    // Suppression de l'index 1 du tableau soit "Février"
    mois = append(mois[:indexASupprimer], mois[(indexASupprimer + 1):]...)

    fmt.Println(mois)
}
