package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel untuk menyimpan input user
	var input int

	// Menampilkan prompt untuk input user
	fmt.Print("Masukkan batas bilangan ganjil: ")

	// Membaca input user
	fmt.Scanln(&input)

	// Deklarasi variabel untuk menyimpan jumlah bilangan ganjil
	var jumlah int

	// For loop untuk menghitung jumlah bilangan ganjil
	for i := 1; i <= input; i++ {
		// Cek apakah bilangan ganjil
		if i%2 != 0 {
			// Menampilkan bilangan ganjil
			fmt.Print(i, " ")
			// Jika ganjil, tambahkan ke jumlah bilangan ganjil
			jumlah++
		}
	}

	// Mencetak baris baru
	fmt.Println()

	// Mencetak jumlah bilangan ganjil
	fmt.Println("Jumlah bilangan ganjil dari 1 hingga", input, "adalah", jumlah)
}
