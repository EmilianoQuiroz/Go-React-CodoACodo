//***********************//
//   Hola mundo en Go	 //
//***********************//

// Primero tenemos que definir el paquete main que no permitirá ejecutar el archivo
package main

// Para mostrar datos por pantalla o consola importar el paquete fmt 
import "fmt"

// Para ejecutar nuestro código se tiene que definir la función main y dentro de esta función
func main() {

	//Para mostrar un Hola mundo usamos el paquete fmt y su método Println
	fmt.Println("Hola Go")
}

//******************************************//
// Como correr un progrma desde la terminal //
//******************************************//
//Comando: 
// go run "nombre del programa"
