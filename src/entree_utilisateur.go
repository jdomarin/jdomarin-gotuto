package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)  // création du scanneur
    fmt.Print("Entrez quelque chose : ")
    scanner.Scan()  			   // lancement du scanneur
    entreeUtilisateur := scanner.Text()    // stockage du résultat
    fmt.Println(entreeUtilisateur)
}
