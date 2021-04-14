package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "time"
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

    switch {
    case choix > 2000:
       fmt.Println("Ah un 2000 !")
    case choix >= 1939 && choix <= 1945:
       fmt.Println("Triste année !")
    case time.Now().Weekday() == time.Sunday:
       fmt.Println("Dimanche !")
    default:
       fmt.Println("Mauvais choix !") // Valeur par défaut
    }

}
