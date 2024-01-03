package main

import (
	"errors"
	"fmt"
)

// Función para calcular el promedio de calificaciones
func calcularPromedio(calificaciones ...int) (float64, error) {
	// Verificar si hay calificaciones negativas
	for _, calificacion := range calificaciones {
		if calificacion < 0 {
			return 0, errors.New("No se permiten calificaciones negativas")
		}
	}

	// Calcular el promedio
	suma := 0
	for _, calificacion := range calificaciones {
		suma += calificacion
	}

	promedio := float64(suma) / float64(len(calificaciones))
	return promedio, nil
}

func main() {
	// Ejemplo de uso
	calificacionesJuan := []int{90, 85, 92, 88, 95}
	promedioJuan, err := calcularPromedio(calificacionesJuan...)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("El promedio de Juan es %.2f\n", promedioJuan)
	}

	// Ejemplo con calificación negativa
	calificacionesMaria := []int{75, -80, 88, 92, 90}
	promedioMaria, err := calcularPromedio(calificacionesMaria...)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("El promedio de Maria es %.2f\n", promedioMaria)
	}
}
