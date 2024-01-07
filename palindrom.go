package main

import "fmt"

// Fungsi untuk memeriksa apakah sebuah string palindrom atau tidak
func isPalindrome(input string) bool {
	// Melakukan perulangan melalui setengah pertama dari string
	for i := 0; i < len(input)/2; i++ {
		// Membandingkan karakter dari awal dan akhir string
		if input[i] != input[len(input)-i-1] {
			// Jika karakter tidak sama, maka akan di return false (bukan palindrom)
			return false
		}
	}
	// Jika perulangan selesai tanpa mereturn false, string tersebut adalah palindrom, sehingga mereturn true
	return true
}

func main() {
	// Meminta user meng-input sebuah string
	fmt.Print("Masukkan sebuah string: ")
	var str string
	fmt.Scanf("%s", &str)

	// Memanggil fungsi isPalindrome dan menyimpan hasilnya dalam variabel 'result'
	result := isPalindrome(str)

	// Memeriksa hasil dan mencetak pesan yang sesuai
	if result == true {
		fmt.Println("String tersebut adalah Palindrom")
	} else {
		fmt.Println("String tersebut bukan Palindrom")
	}
}
