package main

import (
    "fmt"
    "time"
)

func main() {
    switch time.Now().Weekday() {
    case time.Saturday:
       fmt.Println("Il est samedi")
    case time.Sunday:
       fmt.Println("Il est dimanche")
    default:
       fmt.Println("au boulot ! Le week-end est fini !")
    }

}
