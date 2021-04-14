package main

import "fmt"

func main() {
    var a int = 20
    var ap *int
    ap = &a

    fmt.Printf("Valeur de la variable a : %d\n", a)

    *ap = 30 // modification de la valeur de la variable a depuis un pointeur

    fmt.Printf("Valeur de la variable a : %d\n", a)
}
