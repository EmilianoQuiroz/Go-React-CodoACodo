/*******************/
/* Funciones en Go */
/*******************/

package main

import "fmt"

/*
* En Go podemos crear funciones por fuera de nuestra funcion principal y llamarlas dentro de esta
*/
func saludar(nombre string) {
	//fmt.Println("Hola Go")
	fmt.Println("Hola ", nombre)
}

/*
* Ejemplo de funcion que retorna un valor
*/
func sumar(numero1, numero2 int) int{
	return numero1 + numero2
}

func main(){

	// Llamamos a nuestra funcion saludar dentro de la funcion main
	saludar("Santiago")

	// Guardamos el resultado de nuestra funcion sumar en la variable resultado 
	resultado := sumar(23,27)
	//Luego imprimimos la variable resultado por pantalla
	fmt.Println(resultado)

}