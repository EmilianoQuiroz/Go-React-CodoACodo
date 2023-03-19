//******************//
// Constantes en Go //
//******************//
package main
import "fmt"

// Los contrastes se definen como variables y almacenan datos pero los constantes el valor
// asignado no se modifican, los constantes se puede definir con un tipo de dato o también
// sin un tipo de dato solo se le asigna un valor y de acuerdo a su valor se convierte en tipo
// de constante.

//Definir una constante
const s string = "constante"

func main() {

	const n = 25 //Constante de tipo entero

	fmt.Println(s , n)

}