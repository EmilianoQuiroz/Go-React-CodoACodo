/****************/
/* Arrays en GO */
/****************/

package main

import "fmt"

func main(){

	// Arrays en GO
	// Para crear un array debemos definir la cantidad de elementos que tendra 
	// y tambien el tipo de dato de estos elementos 
	var numeros [5]int

	// Podemos asignarle un valor a cada elemento indicando el indice donde lo vamos a guardar
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40

	// Imprimimos en pantalla el array con los datos introducidos
	fmt.Println(numeros)
	// Tambien podemos imprimir un indice especifico 
	fmt.Println(numeros[3])


	// Array con valores
	nombres:= [3]string{"Juan", "pedro", "Luis"}

	fmt.Println(nombres)

	// Tambien podemos crear arrays sin definir la cantidad de elementos que tendra
	colores:= [...]string{
		"Rojo",
		"Azul",
		"Amarillo",
	}

	// Con "len" obtendremos la cantidad de elementos que contiene nuestro array
	fmt.Println(colores, len(colores))
}