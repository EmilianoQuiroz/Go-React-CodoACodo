/****************/
/* Slicen en Go */
/****************/

package main

import "fmt"

func main(){
	// La principal diferencia entre un slicen y un arreglo es que este es mutable
	// es decir, podemos cambiar la longitud de este

	numeros:= []int{1, 2, 3}

	// Imprimimos nuestro slice y la cantidad de elementos que este contiene
	fmt.Println(numeros, len(numeros))

	// Agregamos elementos a nuestro slice
	// Con append podemos agregar elementos a nuestro slicen 
	// Este recibe dos parametros, el primero es el slice a modificar,
	// el segundo parametro es el elemento a agregar
	numeros = append(numeros,4)
	fmt.Println(numeros, len(numeros))

}
