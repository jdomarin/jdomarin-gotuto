package main

import (
   "bufio"
   "fmt"
   "net"
   "os"
   "time"
)

func gestionErreur(err error) {
    if err != nil {
        panic(err)
    }
}

func writeLog(text string, file *os.File) {
    _, err := file.WriteString(text)
    gestionErreur(err)
}


const (
    IP = "127.0.0.1" // IP local
    PORT = "3569"    // Port utilisé
)

func read(conn net.Conn) {
    message, err := bufio.NewReader(conn).ReadString('\n')
    gestionErreur(err)

    fmt.Print("Client : ", string(message))
}

func main() {

    fmt.Println("Lancement du serveur ...")
    file, err := os.OpenFile("chat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
    defer file.Close() // on ferme automatiquement à la fin de notre programme

    ln, err := net.Listen("tcp", fmt.Sprintf("%s:%s", IP, PORT))
    gestionErreur(err)

    var clients []net.Conn // tableau de clients

    for {
        conn, err := ln.Accept()
        if err == nil {
            clients = append(clients, conn) // quand un client se connecte on le rajoute à notre tableau
        }
        gestionErreur(err)
        adresse := conn.RemoteAddr()

        fmt.Println("Un client est connecté depuis", adresse)
        writeLog(fmt.Sprintf("Connexion depuis %s à %s\n", adresse, time.Now()),
                 file)

        go func() { // création de notre goroutine quand un client est connecté
            buf := bufio.NewReader(conn)

            for {
                name, err := buf.ReadString('\n')

                if err != nil {
                    fmt.Printf("Le client s'est déconnecté.\n")
                    writeLog(fmt.Sprintf("Déconnexion depuis %s à %s\n", adresse, time.Now()),
                             file)
                    break
                }
                for _, c := range clients {
                    c.Write([]byte(name)) // on envoie un message à chaque client
                }
            }
        }()
    }
}
