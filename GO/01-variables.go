//*****************//
// Variables en Go //
//*****************//

// Primero tenemos que definir el paquete main que no permitirá ejecutar el archivo
package main

// Para mostrar datos por pantalla o consola importar el paquete fmt 
import "fmt"

// Para ejecutar nuestro código se tiene que definir la función main y dentro de esta función
func main() {

	//Definir una variable
	var nombre string
	nombre = "Santiago" //Asignar un Valor
	
	fmt.Println(nombre) //Mostrar el valor de la Variable
	nombre = "Quiroz" //Modificar el valor de la Variable
	fmt.Println(nombre)

	//En Go podemos definir variables de muchas formas, como definir múltiples variables de un solo tipo. 
	//Definir variables múltiples de un solo tipo
	var a, b int = 1, 2
	var c, d int
	//Asignado valores
	c = 3
	d = 4

	//Mostrar por pantalla
	fmt.Println(a, b, c, d)

	// También definir múltiples variables de varios tipos de datos. 
	//Definir variables múltiples de varios tipos
var (
	pi float64
	booleano bool
	cadena = "texto 01"
	edad = 25
	)
	pi = 3.141592
	booleano = true
	fmt.Println(pi, booleano, cadena, edad)
	
	//Definir variable sin tipo de dato
	v1 := 24 // Int
	v2 := "Texto 02" // string
	fmt.Println(v1, v2)	

}
