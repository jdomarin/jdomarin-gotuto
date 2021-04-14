package main

import "fmt"

func main() {
    notes := map[string]int{"Hatim": 20, "Alex": 18, "Kevin": 15, "Robert": 17}
    for eleve := range notes {
        fmt.Println("La note de ", eleve, "est", notes[eleve])
    }

    fmt.Println(notes)

    delete(notes, "Hatim")
    fmt.Println(notes)
}
