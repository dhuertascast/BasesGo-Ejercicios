package alimento

import (
	"fmt"
)

// calcularCantidadAlimento calcula la cantidad total de alimento necesario para un tipo de animal.
func calcularCantidadAlimento(tipoAnimal string, cantidad float64) float64 {
	var factorConversion float64

	switch tipoAnimal {
	case "dog":
		factorConversion = 10.0
	case "cat":
		factorConversion = 5.0
	case "hamster":
		factorConversion = 0.25
	case "tarantula":
		factorConversion = 0.15
	default:
		return 0
	}

	return cantidad * factorConversion
}

func main() {
	fmt.Println(calcularCantidadAlimento("dog", 10))
}
