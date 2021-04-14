package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

const (
    lignes = 3
    colonnes = 3
)

var grille = [lignes][colonnes]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}

func afficherGrille() {
    for i, ligne := range grille {
        for j, _ := range ligne {
            fmt.Print(grille[i][j])
            fmt.Print(" ")
        }
        fmt.Println("")
    }
}

func jouer(nombre int, symbole string) bool {
    for i, ligne := range grille {
        for j, colonne := range ligne {
            valeur, ok := strconv.Atoi(colonne)
            if ok == nil && nombre == valeur {
                grille[i][j] = symbole
                return true
            }
        }
    }
    return false
}

func finirPartie(joueur int, symbole string) bool {
    // diagonale avec un même symbole
    if (symbole == grille[0][0] && symbole == grille[1][1] && symbole == grille[2][2]) || (symbole == grille[0][2] && symbole == grille[1][1] && symbole == grille[2][0]) {
        fmt.Println("Joueur", joueur, "vous avez gagné !")
        return true
    }
    for i, ligne := range grille {
        // ligne avec un même symbole
        if symbole == grille[i][0] && symbole == grille[i][1] && symbole == grille[i][2] {
            fmt.Println("Joueur", joueur, "vous avez gagné !")
            return true
        }
        for j, _ := range ligne {
            // colonne avec un même symbole
            if symbole == grille[0][j] && symbole == grille[1][j] && symbole == grille[2][j] {
                fmt.Println("Joueur", joueur, "vous avez gagné !")
                return true
            }
        }
    }
    return false
}

func main() {
    i := 1 // joueur 1 commence
    symbole := "X" // joueur 1 a le symbole X
    cases_disponibles := lignes * colonnes
    for { // commence la boucle
        afficherGrille()
        fmt.Printf("Joueur %d, entrez un nombre compris entre 1 et 9: ", i)
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        choix, ok := strconv.Atoi(scanner.Text())
        if ok != nil {
            fmt.Println("Entrez un nombre et non autre chose !")
            continue
        }
        if choix < 1 || choix > 9 {
            fmt.Println("Votre nombre doit être compris entre 1 et 9 !")
            continue
        }
        jeu := jouer(choix, symbole)
        if !jeu {
            fmt.Println("Cette case est déjà prise !")
            continue
        } else {
            cases_disponibles-- // enlève une case disponible
            if cases_disponibles == 0 { // la grille est remplie
                afficherGrille()
                fmt.Println("Partie nulle !")
                break  // le jeu est terminé
            }
        }
        fini := finirPartie(i, symbole)
        if fini { // un des deux joueurs a gagné
            break  // le jeu est terminé
        } else { // le jeu continue
            if i == 1 { // joueur 1
                i++  // passer au joueur 2 avec le symbole O
                symbole = "O"
            } else { // joueur 2
                i--  // revenir au joueur 1 avec le symbole X
                symbole = "X"
            }
        } // finirPartie
    } // for
} // main

