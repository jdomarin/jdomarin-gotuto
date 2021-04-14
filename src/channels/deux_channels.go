package main

import (
   "fmt"
   "time"
)

func run(ch chan string, name string) {
    time.Sleep(time.Second * 2)
    fmt.Println("fonction run() :", name)
    ch <- name
}

func main() {

    now := time.Now()
    ch := make(chan string)

    go run(ch, "channel 1")
    go run(ch, "channel 2")

    // changement d'ordre de lecture de nos channels
    fmt.Println("fonction main() :", <-ch)
    fmt.Println("fonction main() :", <-ch)

    fmt.Println(time.Now().Sub(now))
}
