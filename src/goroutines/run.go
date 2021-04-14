package main

import (
    "fmt"
    "time"
)

func run(name string) {
    for i:= 0; i < 3; i++ {
        time.Sleep(1 * time.Second) // attendre 1 seconde
        fmt.Println(name, " : ", i)
    }
}

func main() {
    debut := time.Now()
    go run("Hatim")
    go run("Robert")
    run("Alex")
    fin := time.Now()
    fmt.Println(fin.Sub(debut))
}
