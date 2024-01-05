// archivo: promedio_test.go

package impuesto_test

import "testing"

func TestCalcularPromedioVacio(t *testing.T) {
	notas := []float64{}
	promedio := calcularPromedio(notas)

	// Verificar que el promedio de un slice vacío sea 0
	if promedio != 0 {
		t.Errorf("El promedio de un slice vacío debería ser 0, pero es %v", promedio)
	}
}

func TestCalcularPromedioUnaNota(t *testing.T) {
	notas := []float64{85.5}
	promedio := calcularPromedio(notas)

	// Verificar que el promedio de un slice con una nota sea igual a esa nota
	if promedio != notas[0] {
		t.Errorf("El promedio de un slice con una nota debería ser %v, pero es %v", notas[0], promedio)
	}
}

func TestCalcularPromedioMultipleNotas(t *testing.T) {
	notas := []float64{90.0, 85.0, 95.0, 88.0}
	promedio := calcularPromedio(notas)

	// Verificar que el promedio sea calculado correctamente
	promedioEsperado := (90.0 + 85.0 + 95.0 + 88.0) / 4.0
	if promedio != promedioEsperado {
		t.Errorf("El promedio debería ser %v, pero es %v", promedioEsperado, promedio)
	}
}
