// archivo: cantidad_alimento_test.go

package alimento_test

import "testing"

func TestCalcularCantidadAlimentoPerro(t *testing.T) {
	cantidad := 10.0
	alimento := calcularCantidadAlimento("dog", cantidad)

	// Verificar que la cantidad de alimento para perros sea calculada correctamente
	alimentoEsperado := cantidad * 10.0
	if alimento != alimentoEsperado {
		t.Errorf("La cantidad de alimento para perros debería ser %v, pero es %v", alimentoEsperado, alimento)
	}
}

func TestCalcularCantidadAlimentoGato(t *testing.T) {
	cantidad := 5.0
	alimento := calcularCantidadAlimento("cat", cantidad)

	// Verificar que la cantidad de alimento para gatos sea calculada correctamente
	alimentoEsperado := cantidad * 5.0
	if alimento != alimentoEsperado {
		t.Errorf("La cantidad de alimento para gatos debería ser %v, pero es %v", alimentoEsperado, alimento)
	}
}

func TestCalcularCantidadAlimentoHamster(t *testing.T) {
	cantidad := 4.0
	alimento := calcularCantidadAlimento("hamster", cantidad)

	// Verificar que la cantidad de alimento para hamsters sea calculada correctamente
	alimentoEsperado := cantidad * 0.25
	if alimento != alimentoEsperado {
		t.Errorf("La cantidad de alimento para hamsters debería ser %v, pero es %v", alimentoEsperado, alimento)
	}
}

func TestCalcularCantidadAlimentoTarantula(t *testing.T) {
	cantidad := 2.0
	alimento := calcularCantidadAlimento("tarantula", cantidad)

	// Verificar que la cantidad de alimento para tarántulas sea calculada correctamente
	alimentoEsperado := cantidad * 0.15
	if alimento != alimentoEsperado {
		t.Errorf("La cantidad de alimento para tarántulas debería ser %v, pero es %v", alimentoEsperado, alimento)
	}
}
