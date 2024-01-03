package main

import "fmt"

// Función para calcular el impuesto sobre el salario
func calcularImpuesto(salario float64) float64 {
	var impuesto float64

	if salario > 150000 {
		// Si el salario es mayor a $150,000, descuento del 27%
		impuesto = salario * 0.27
	} else if salario > 50000 {
		// Si el salario es mayor a $50,000 pero no mayor a $150,000, descuento del 17%
		impuesto = salario * 0.17
	}

	return impuesto
}

func main() {
	// Ejemplo de uso
	salarioJuan := 60000.0
	impuestoJuan := calcularImpuesto(salarioJuan)

	fmt.Printf("Juan gana $%.2f y su impuesto es $%.2f\n", salarioJuan, impuestoJuan)
}
