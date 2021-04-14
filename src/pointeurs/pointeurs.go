package main

import "fmt"

func main() {
    var a int = 20
    var ap *int // création d'un pointeur
    ap = &a // stockage de l'adresse mémoire de la variable dans notre pointeur

    fmt.Printf("Adresse de votre variable a : %p\n", &a)

    fmt.Printf("Valeur de votre variable (pointeur) ap : %p\n", ap)

    fmt.Printf("Valeur de l'adresse %p: %d\n", ap, *ap)
}
