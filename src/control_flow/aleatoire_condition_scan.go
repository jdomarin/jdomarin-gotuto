package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "time"
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

    fmt.Print("Entrez votre prénom : ")
    scanner.Scan()
    prenom := scanner.Text()

    /*
        On a besoin de changer la graine (générateur de nombres pseudo-aléatoires) à
        chaque exécution de programme sinon on obtiendra le même nombre aléatoire
    */
    rand.Seed(time.Now().UnixNano())
    radomInt := rand.Intn(2) // retourne aléatoirement soit 1 soit 0
    radomInt2 := rand.Intn(2)

    if age < 18 {
        fmt.Println("Sortez !")
    } else if prenom == "Hatim" || prenom == "hatim" {
        fmt.Println("Ah c'est toi Hatim, dehors !")
    } else if age == 18 && radomInt == 1 { // si le client est chanceux et qu'il a 18 ans
        fmt.Println("Hum vous avez de la chance, je suis de bonne humeur, entrez !")
    } else if radomInt2 == 0 {
        fmt.Println("Vous n'avez vraiment pas de chance, sortez !")
    } else {
        fmt.Println("Entrez :)")
    }

}
