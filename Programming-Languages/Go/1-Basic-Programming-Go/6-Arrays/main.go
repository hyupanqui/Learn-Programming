package main

import "fmt"

func main() {
	// Declaraciรณn de un array
	var food [3]string

	// Endicamos el indice donde queremos almacenar el valor
	food[0] = "๐"
	food[1] = "๐"
	food[2] = "๐ญ"
	fmt.Println("Comidas: ", food)

	// Array literales
	fruit := [3]string{"๐", "๐", "๐"}
	fmt.Println("Frutas: ", fruit)

	// Array son imutables
	var array [4]int

	array[0] = 1
	array[1] = 2

	fmt.Println(array, len(array), cap(array))
}
