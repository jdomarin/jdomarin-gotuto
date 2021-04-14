package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Votre choix : ")
    scanner.Scan()
    choix, err := strconv.Atoi(scanner.Text())
    if err != nil {
        fmt.Println("Entrez un entier !")
        os.Exit(2)
    }

    switch choix { // la variable que l'on souhaite vérifier
        case 0, 1: // 1 ou 0
            fmt.Println("George Boole !")
        case 7:
            fmt.Println("William van de Walle !")
        case 23:
            fmt.Println("Pour certains le nombre 23 est source de nombreux mystères, qu'en penses-tu Jim Carrey?")
        case 42:
            fmt.Println("la réponse ultime à la question du sens de la vie")
        case 666:
            fmt.Println("Quand le diable veut une âme, le mal devient séduisant")
        default:
            fmt.Println("Mauvais choix !") // Valeur par défaut
    }
}
