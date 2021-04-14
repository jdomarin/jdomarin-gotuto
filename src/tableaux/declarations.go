package main

import "fmt"

func main() {
    var tableauInt [10]int
    var tableauFloat [10]float32
    var tableauString [10]string
    var tableauBool[10] bool
    fmt.Println("Valeur par défaut de la variable tableauInt :", tableauInt)
    fmt.Println("Valeur par défaut de la variable tableauFloat :", tableauFloat)
    fmt.Println("Valeur par défaut de la variable tableauString :", tableauString)
    fmt.Println("Valeur par défaut de la variable tableauBool :", tableauBool)

   var tableau1 = [5]int{1, 2, 3, 4, 5}
   fmt.Println("la taille de mon tableau1 :", len(tableau1))
   fmt.Println("les valeurs de mon tableau1 :", tableau1)
}
