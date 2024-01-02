package main

import "fmt"

func main() {
	// Variable de entrada: número del mes
	fmt.Println("Ingresa el numero del mes")
	var numeroMes int
	fmt.Scanln(&numeroMes)
	// Validar el número del mes
	if numeroMes >= 1 && numeroMes <= 12 {
		// Mapeo de números de mes a nombres de mes
		meses := map[int]string{
			1:  "Enero",
			2:  "Febrero",
			3:  "Marzo",
			4:  "Abril",
			5:  "Mayo",
			6:  "Junio",
			7:  "Julio",
			8:  "Agosto",
			9:  "Septiembre",
			10: "Octubre",
			11: "Noviembre",
			12: "Diciembre",
		}

		// Imprimir el mes correspondiente utilizando el mapa
		fmt.Println(meses[numeroMes])
	} else {
		fmt.Println("Número de mes incorrecto. Debe estar entre 1 y 12.")
	}
}
