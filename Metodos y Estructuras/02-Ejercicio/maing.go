package main

import (
	"fmt"
	"time"
)

// Person es una estructura que representa a una persona.
type Person struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}

// Employee es una estructura que representa a un empleado, compuesta por una Person.
type Employee struct {
	ID       int
	Position string
	Person   Person
}

// PrintEmployee imprime los campos de un empleado.
func (e *Employee) PrintEmployee() {
	fmt.Printf("ID: %d, Name: %s, DateOfBirth: %s, Position: %s\n", e.Person.ID, e.Person.Name, e.Person.DateOfBirth.Format("2006-01-02"), e.Position)
}

func main() {
	// Instanciar una Person
	person := Person{
		ID:          1,
		Name:        "John Doe",
		DateOfBirth: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	// Instanciar un Employee con composición de la Person
	employee := Employee{
		ID:       101,
		Position: "Developer",
		Person:   person,
	}

	// Ejecutar el método PrintEmployee
	employee.PrintEmployee()
}
