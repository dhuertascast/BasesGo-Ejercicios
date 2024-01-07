// archivo: salario_test.go

package salario_test

import "testing"

func TestCalcularSalarioCategoriaC(t *testing.T) {
	minutosTrabajados := 1200
	categoria := "C"
	salario := calcularSalarioCategoria(minutosTrabajados, categoria)

	// Verificar que el salario sea calculado correctamente para la categoría "C"
	salarioEsperado := 1000 * float64(minutosTrabajados) / 60
	if salario != salarioEsperado {
		t.Errorf("El salario debería ser %v, pero es %v", salarioEsperado, salario)
	}
}

func TestCalcularSalarioCategoriaB(t *testing.T) {
	minutosTrabajados := 1200
	categoria := "B"
	salario := calcularSalarioCategoria(minutosTrabajados, categoria)

	// Verificar que el salario sea calculado correctamente para la categoría "B"
	salarioEsperado := 1500*(float64(minutosTrabajados)/60) + (1500*(float64(minutosTrabajados)/60))*0.20
	if salario != salarioEsperado {
		t.Errorf("El salario debería ser %v, pero es %v", salarioEsperado, salario)
	}
}

func TestCalcularSalarioCategoriaA(t *testing.T) {
	minutosTrabajados := 1200
	categoria := "A"
	salario := calcularSalarioCategoria(minutosTrabajados, categoria)

	// Verificar que el salario sea calculado correctamente para la categoría "A"
	salarioEsperado := 3000*(float64(minutosTrabajados)/60) + (3000*(float64(minutosTrabajados)/60))*0.50
	if salario != salarioEsperado {
		t.Errorf("El salario debería ser %v, pero es %v", salarioEsperado, salario)
	}
}
