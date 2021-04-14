package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Entrez un nombre entier : ")
    scanner.Scan()

    nbr, _ := strconv.Atoi(scanner.Text()) // conversion du type string en int

    fmt.Printf("res : %d\n", (nbr + 6))
}
