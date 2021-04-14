package main

import "fmt"

func run(c chan string, name string) {
    c <- name // envoyer une valeur d'un channel
}

func main() {
    canal := make(chan string)
    go run(canal, "Hatim")
    fmt.Println(<-canal) // récupérer une valeur d'un channel
}
