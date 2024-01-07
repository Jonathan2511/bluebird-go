package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial dari suatu bilangan bulat positif menggunakan rekursif
func factorial(n int) int {
	// Basis rekursif, jika n sama dengan 0 atau 1, kembalikan 1
	if n == 0 || n == 1 {
		return 1
	}
	// Memanggil fungsi factorial dengan nilai yang lebih kecil dari n dan mengalikan hasilnya dengan n
	return n * factorial(n-1)
}

func main() {
	// Meminta user meng-input angka
	fmt.Print("Masukkan angka untuk menghitung faktorial: ")
	var input int
	fmt.Scan(&input)

	// Validasi apakah input positif
	if input < 0 {
		fmt.Println("Tolong Masukkan bilangan bulat positif!")
		return
	}

	// Menghitung faktorial dan mencetak hasilnya
	result := factorial(input)
	fmt.Println("Faktorial dari", input, "adalah", result)
}
