package main

import (
    "affichage"
    "affichage/passion" // chemin relatif
    "fmt"
)

func main() {
    fmt.Println("Je m'appelle", affichage.Nom)
    fmt.Println(affichage.AfficheSexe())
    fmt.Println(passion.AffichagePassion())
}
