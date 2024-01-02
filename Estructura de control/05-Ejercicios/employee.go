package main

import "fmt"

func main() {
	// Mapa de empleados con sus edades
	employees := map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Darío":    44,
		"Pedro":    30,
	}

	// Imprimir la edad de Benjamin
	edadBenjamin, existe := employees["Benjamin"]
	if existe {
		fmt.Printf("La edad de Benjamin es: %d\n", edadBenjamin)
	} else {
		fmt.Println("Benjamin no se encuentra en la lista de empleados.")
	}

	// Contar empleados mayores de 21 años
	mayoresDe21 := 0
	for _, edad := range employees {
		if edad > 21 {
			mayoresDe21++
		}
	}
	fmt.Printf("Hay %d empleados mayores de 21 años.\n", mayoresDe21)

	// Agregar un nuevo empleado llamado Federico con 25 años
	employees["Federico"] = 25
	fmt.Println("Se ha agregado a Federico a la lista de empleados.")

	// Eliminar a Pedro del mapa
	delete(employees, "Pedro")
	fmt.Println("Pedro ha sido eliminado de la lista de empleados.")

	// Imprimir la lista actualizada de empleados
	fmt.Println("Lista de empleados actualizada:")
	for nombre, edad := range employees {
		fmt.Printf("%s: %d años\n", nombre, edad)
	}
}
