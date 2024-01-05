package inpuesto

import (
	"fmt"
)

// calcularPromedio calcula el promedio de las notas de los alumnos.
func calcularPromedio(notas []float64) float64 {
	if len(notas) == 0 {
		return 0
	}

	suma := 0.0
	for _, nota := range notas {
		suma += nota
	}

	return suma / float64(len(notas))
}

func main() {
	notas := []float64{10, 20, 30, 40, 50}
	fmt.Println(calcularPromedio(notas))
}
