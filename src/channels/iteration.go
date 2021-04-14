package main

import "fmt"

func main() {
    ch := make(chan string, 2) // buffer de taille 2

    go func() {
        defer close(ch) // on indique à notre compilateur qu'on a fini d'écrire sur notre channel
        ch <- "test" // 1 seul récepteur
    }()

    for true {
        if elem, ok := <-ch; ok == true { // est-ce que le channel possède encore un récepteur ?
            fmt.Println(elem)
        } else {
            break
        }
    }
}
