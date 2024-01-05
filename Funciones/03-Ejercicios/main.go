package main

import (
	"errors"
	"fmt"
)

// Función para calcular el salario de un empleado
func calcularSalario(minutosTrabajadosPorMes int, categoria string) (float64, error) {
	// Validar la categoría
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
		return 0, errors.New("Categoría no válida")
	}

	// Calcular salario
	salarioPorHora := tarifaPorHora * float64(minutosTrabajadosPorMes) / 60
	salarioTotal := salarioPorHora + (salarioPorHora * porcentajeAdicional)

	return salarioTotal, nil
}

func main() {
	// Ejemplo de uso
	minutosTrabajados := 1200
	categoriaJuan := "B"
	salarioJuan, err := calcularSalario(minutosTrabajados, categoriaJuan)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("El salario de Juan es $%.2f\n", salarioJuan)
	}

	// Ejemplo con categoría no válida
	categoriaMaria := "D"
	salarioMaria, err := calcularSalario(minutosTrabajados, categoriaMaria)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("El salario de Maria es $%.2f\n", salarioMaria)
	}
}
