package main

import "fmt"

func main() {
    var jours = [5]string{"Hatim", "Robert", "Inconnu", "Ahmed", "Inconnu"}

    jours[0] = "Alex" // on remplace le premier élément (ici Hatim) par Alex
    fmt.Println(jours)

    jours = replaceByHatim(jours)
    fmt.Println(jours)
}

/*
    J'utilise une fonction pour vous montrer qu'il est
    possible de prendre en paramètre un tableau
    mais aussi de retourner un tableau dans une fonction
*/

func replaceByHatim(jours [5]string) [5]string {
    for i, jour := range jours {
        if jour == "Inconnu" {
            jours[i] = "Hatim" // Remplacer "Inconnu" par "Hatim"
        }
    }
   return jours
}
