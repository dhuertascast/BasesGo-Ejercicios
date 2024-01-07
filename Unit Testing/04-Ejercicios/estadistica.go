package estadisticas

import (
	"fmt"
)

// calcularEstadisticas calcula el mínimo, máximo y promedio de calificaciones.
func calcularEstadisticas(calificaciones []float64) (float64, float64, float64) {
	if len(calificaciones) == 0 {
		return 0, 0, 0
	}

	// Inicializar min y max con el primer elemento
	min := calificaciones[0]
	max := calificaciones[0]
	suma := calificaciones[0]

	// Iterar sobre el resto de elementos para encontrar min, max y calcular la suma
	for _, calificacion := range calificaciones[1:] {
		if calificacion < min {
			min = calificacion
		}
		if calificacion > max {
			max = calificacion
		}
		suma += calificacion
	}

	// Calcular el promedio
	promedio := suma / float64(len(calificaciones))

	return min, max, promedio
}

func main() {
	calificaciones := []float64{10, 20, 30, 40, 50}
	fmt.Println(calcularEstadisticas(calificaciones))
}
