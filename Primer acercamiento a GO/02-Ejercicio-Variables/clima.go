package main

import "fmt"

func main() {
	var temperature float64
	var humidity float64
	var pressure float64

	temperature = 25.5
	humidity = 60.0
	pressure = 1013.25

	fmt.Printf("Temperatura: %.2f °C\n", temperature)
	fmt.Printf("Humedad: %.2f%%\n", humidity)
	fmt.Printf("Presión: %.2f hPa\n", pressure)
}
