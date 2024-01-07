// archivo: estadisticas_test.go

package estadisticas_test

import "testing"

func TestCalcularMinimo(t *testing.T) {
	calificaciones := []float64{90.0, 85.0, 95.0, 88.0}
	minimo := calcularEstadisticas(calificaciones)

	// Verificar que el mínimo sea calculado correctamente
	minimoEsperado := 85.0
	if minimo != minimoEsperado {
		t.Errorf("El mínimo debería ser %v, pero es %v", minimoEsperado, minimo)
	}
}

func TestCalcularMaximo(t *testing.T) {
	calificaciones := []float64{90.0, 85.0, 95.0, 88.0}
	_, maximo, _ := calcularEstadisticas(calificaciones)

	// Verificar que el máximo sea calculado correctamente
	maximoEsperado := 95.0
	if maximo != maximoEsperado {
		t.Errorf("El máximo debería ser %v, pero es %v", maximoEsperado, maximo)
	}
}

func TestCalcularPromedio(t *testing.T) {
	calificaciones := []float64{90.0, 85.0, 95.0, 88.0}
	_, _, promedio := calcularEstadisticas(calificaciones)

	// Verificar que el promedio sea calculado correctamente
	promedioEsperado := (90.0 + 85.0 + 95.0 + 88.0) / 4.0
	if promedio != promedioEsperado {
		t.Errorf("El promedio debería ser %v, pero es %v", promedioEsperado, promedio)
	}
}
