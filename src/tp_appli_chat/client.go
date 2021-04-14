package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "sync"
)

func gestionErreur(err error) {
    if err != nil {
        panic(err)
    }
}

const (
    IP = "127.0.0.1" // IP local
    PORT = "3569"    // Port utilisé
)

func main() {

    var wg sync.WaitGroup
    var pseudo string
    var pseudos []string

    // Connexion au serveur
    conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", IP, PORT))
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Println("Entrez un pseudo (max 20 caractères) :")
        scanner.Scan()
        pseudo = scanner.Text()
        for _, elem := range pseudos {
            if elem == pseudo {
                fmt.Sprintf("Le pseudo %s est déjà pris!\n", pseudo)
                continue
            }
        }
        if 0 < len(pseudo) && len(pseudo) <= 20 {
            fmt.Sprintf("%s entre dans le chat\n")
            pseudos = append(pseudos, pseudo)
            break
        } else {
            fmt.Sprintf("Le pseudo %s fait plus de 20 caractères !\n", pseudo)
        }
    }

    
    gestionErreur(err)

    wg.Add(2)

    go func() { // goroutine dédiée à l'entrée utilisateur
        defer wg.Done()
        for {
            reader := bufio.NewReader(os.Stdin)
            text, err := reader.ReadString('\n')
            gestionErreur(err)

            conn.Write([]byte(text))
        }
    }()

    go func() { // goroutine dédiée à la réception des messages du serveur
        defer wg.Done()
        for {
            message, err := bufio.NewReader(conn).ReadString('\n')
            gestionErreur(err)

            fmt.Print(fmt.Sprintf("serveur (%s) : %s", pseudo, message))
        }
    }()

    wg.Wait()
}
