/*********************/
/* Paquete fmt en Go */
/*********************/
package main

/**
* El paquete fmt sirve para la entrada y salida de datos por pantalla
 */
import "fmt"

func main() {

	//Impresion simple

	hola := "Hola"
	mundo := " mundo"

	// Con Print podemos iprimir los datos en una sola linea
	fmt.Print(hola)
	fmt.Print(mundo)

	nombre := "Santiago"
	edad := 24

	//Com Printf podemos formatear informacion
	fmt.Printf("Hola %s de %d años \n", nombre, edad)

	
}
