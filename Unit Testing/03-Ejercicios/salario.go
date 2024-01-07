package salario

import (
	"fmt"
)

// calcularSalarioCategoria calcula el salario según la categoría.
func calcularSalarioCategoria(minutosTrabajados int, categoria string) float64 {
	var tarifaPorHora float64
	var porcentajeAdicional float64

	switch categoria {
	case "C":
		tarifaPorHora = 1000
		porcentajeAdicional = 0
	case "B":
		tarifaPorHora = 1500
		porcentajeAdicional = 0.20
	case "A":
		tarifaPorHora = 3000
		porcentajeAdicional = 0.50
	default:
		return 0
	}

	salarioPorHora := tarifaPorHora * float64(minutosTrabajados) / 60
	salarioTotal := salarioPorHora + (salarioPorHora * porcentajeAdicional)

	return salarioTotal
}

func main() {
	fmt.Println(calcularSalarioCategoria(100, "C"))
}
