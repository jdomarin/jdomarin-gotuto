package main

import "fmt"

var g int  // déclaration de notre variable formelle

func test(g int) { // déclaration de notre paramètre global
    g += 20
    fmt.Println("Pendant ma fonction test() : ", g)
}

func main() {
    fmt.Println("Avant l'utilisation de la fonction test() :", g)
    test(20)
    fmt.Println("Pendant ma fonction main() :", g)
    g += 30
    fmt.Println("Modifie moi encore : ", g)
}
