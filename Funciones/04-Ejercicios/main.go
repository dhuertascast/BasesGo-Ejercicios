package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

// Función principal que devuelve otra función de cálculo y un mensaje de error si es necesario
func operation(operationType string) (func(...int) float64, error) {
	switch operationType {
	case minimum:
		return minFunction, nil
	case average:
		return averageFunction, nil
	case maximum:
		return maxFunction, nil
	default:
		return nil, errors.New("Operación no válida")
	}
}

// Función para calcular el mínimo de una lista de enteros
func minFunction(numbers ...int) float64 {
	minValue := float64(numbers[0])
	for _, num := range numbers {
		if float64(num) < minValue {
			minValue = float64(num)
		}
	}
	return minValue
}

// Función para calcular el promedio de una lista de enteros
func averageFunction(numbers ...int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

// Función para calcular el máximo de una lista de enteros
func maxFunction(numbers ...int) float64 {
	maxValue := float64(numbers[0])
	for _, num := range numbers {
		if float64(num) > maxValue {
			maxValue = float64(num)
		}
	}
	return maxValue
}

func main() {
	// Ejemplo de uso
	minFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Utilizar las funciones para realizar cálculos
	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	// Imprimir resultados
	fmt.Printf("Mínimo: %.2f\n", minValue)
	fmt.Printf("Promedio: %.2f\n", averageValue)
	fmt.Printf("Máximo: %.2f\n", maxValue)
}
