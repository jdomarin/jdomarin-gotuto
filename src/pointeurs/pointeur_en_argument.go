package main

import "fmt"

func main() {
    var a int = 20

    fmt.Printf("Avant de modifier la variable de a : %d\n", a)

    rajouterCinq(&a)

    fmt.Printf("Après modification par des pointeurs de la variable de a : %d\n", a)

    affichage(&a)
}

// fonction qui prend en paramètre un pointeur
func rajouterCinq(pa *int) {
    *pa = *pa + 5 // modification de la valeur de la variable "a" depuis un pointeur
}

// fonction qui prend en paramètre un pointeur
func affichage(pa *int) {
    fmt.Println("affichage de valeur dans une fonction :", *pa)
}
