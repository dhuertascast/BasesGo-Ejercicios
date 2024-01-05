package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

// Función principal que devuelve otra función de cálculo y un mensaje de error si es necesario
func animal(animalType string) (func(float64) float64, error) {
	switch animalType {
	case dog:
		return dogFunction, nil
	case cat:
		return catFunction, nil
	case hamster:
		return hamsterFunction, nil
	case tarantula:
		return tarantulaFunction, nil
	default:
		return nil, errors.New("Animal no válido")
	}
}

// Función para calcular la cantidad de alimento para un perro
func dogFunction(cantidad float64) float64 {
	return cantidad * 10
}

// Función para calcular la cantidad de alimento para un gato
func catFunction(cantidad float64) float64 {
	return cantidad * 5
}

// Función para calcular la cantidad de alimento para un hamster
func hamsterFunction(cantidad float64) float64 {
	return cantidad * 0.25
}

// Función para calcular la cantidad de alimento para una tarántula
func tarantulaFunction(cantidad float64) float64 {
	return cantidad * 0.15
}

func main() {
	// Ejemplo de uso
	animalDog, msg := animal(dog)
	animalCat, msg := animal(cat)

	if msg != nil {
		fmt.Println("Error:", msg)
		return
	}

	// Utilizar las funciones para calcular la cantidad de alimento
	var cantidad float64

	cantidad += animalDog(10)
	cantidad += animalCat(5)

	fmt.Printf("Cantidad total de alimento: %.2f kg\n", cantidad)
}
