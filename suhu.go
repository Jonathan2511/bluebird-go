package main

import (
	"fmt"
)

// Fungsi untuk mengkonversi suhu dari Celsius ke Fahrenheit
func convertCelsiusToFahrenheit(celsius float64) float64 {
	// Rumus konversi
	fahrenheit := (celsius * 9 / 5) + 32

	// Return hasil konversi
	return fahrenheit
}

func main() {
	// Meminta user input suhu dalam Celsius
	fmt.Print("Masukkan suhu dalam Celsius: ")
	var celsius float64
	fmt.Scanf("%f", &celsius)

	// Konversi suhu ke Fahrenheit
	fahrenheit := convertCelsiusToFahrenheit(celsius)

	// Mencetak suhu dalam Fahrenheit
	fmt.Println("Suhu dalam Fahrenheit:", fahrenheit)
}
