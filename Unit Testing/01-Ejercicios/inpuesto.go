package inpuesto

import (
	"fmt"
)

func impuestoSalario(salario float64) float64 {
	var impuesto float64

	if salario > 150000 {
		impuesto = salario * 0.27
	} else if salario > 50000 {
		impuesto = salario * 0.17
	}
	fmt.Println(impuesto)
	return impuesto
}

func main() {
	impuestoSalario(100000)
}
