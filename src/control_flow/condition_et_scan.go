package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Entrez votre age : ")
    scanner.Scan()

    age, err := strconv.Atoi(scanner.Text())
    if err != nil {
        // Si la conversion n'a pas fonctionné alors on affiche un message d'erreur et on quitte le programme
        fmt.Println("On essaie de m'arnaquer ? allé Dehors ! Et la prochaine fois entrez un entier !")
        os.Exit(2)  // on quitte le programme
    }

    if age < 18 {
        fmt.Println("Sortez !")
    } else {
        fmt.Println("Entrez !")
    }

}
