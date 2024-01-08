package main

import (
	"fmt"
	"time"
)

// Alumno es una estructura que representa a un estudiante.
type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    time.Time
}

// detalle imprime el detalle de los datos de un alumno.
func (a *Alumno) detalle() {
	fmt.Printf("Name: %s\nApellido: %s\nDNI: %s\nFecha: %s\n", a.Nombre, a.Apellido, a.DNI, a.Fecha.Format("2006-01-02"))
}

func main() {
	// Instanciar un alumno
	alumno := Alumno{
		Nombre:   "Juan",
		Apellido: "Perez",
		DNI:      "12345678",
		Fecha:    time.Now(),
	}

	// Ejecutar el método detalle para imprimir los datos del alumno
	alumno.detalle()
}
