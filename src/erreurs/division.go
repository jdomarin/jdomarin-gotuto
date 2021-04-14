package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func division() {
    for true {
        scanner := bufio.NewScanner(os.Stdin)
        fmt.Print("Entrez un chiffre : ")
        scanner.Scan()
        nbr, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Println("Vous devez rentrer un nombre et non une chaîne de caractères !")
        } else if nbr <= 0 {
            fmt.Println("[division par zéro impossible] Votre valeur doit être supérieure à 0")
        } else {
            fmt.Println("Résultat :", 1000/nbr)
            break
        }
    }
}

func main() {
   division()
   fmt.Println("Fin")
}
