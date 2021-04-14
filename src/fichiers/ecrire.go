package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func write(text string, file *os.File) {
    if _, err := file.WriteString(text); err != nil {
        panic(err)
    }
}

func read(filename string) string {
    data, err := ioutil.ReadFile(filename) // lire le fichier
    check(err)
    return string(data)
}

func main() {
    file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
    defer file.Close() // on ferme automatiquement à la fin de notre programme
    check(err)

    write("Test\n", file)

    data := read(file.Name())
    fmt.Println(data)
}
