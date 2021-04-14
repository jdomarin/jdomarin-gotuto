package main

import "fmt"

func main() {

    animaux1 := []string{"Lion", "Cheval", "Ours"}
    fmt.Println("Contenu du tableau animaux1 :", animaux1)

    // création d'une slice cible avec la même taille que la slice source
    animaux2 := make([]string, len(animaux1))

    // copie du contenu de la slice source vers la slice cible
    copy(animaux2, animaux1)

    fmt.Println("Contenu du tableau animaux2 :", animaux2)
}
