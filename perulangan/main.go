package main

import "fmt"

func main() {
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println("Angka", i)
	// }

	/*
		for (hanya kondisi)
		Konsepnya mirip seperti while milik bahasa pemrograman lain.
	*/
	// x := 1

	// for x < 5 {
	// 	fmt.Println("Perulangan Konsep while : ", x)
	// 	x++
	// }

	/*
		perulangan dengan menggunakan kombinasi keyword for dan range.
		Cara ini biasa digunakan untuk me-looping data bertipe array
	*/

	title := "Golang the best language"

	// for index, letter := range title {
	// 	fmt.Println("Index : ", index, "Letter : ", string(letter))
	// }

	// tanpa variable index
	for _, letter := range title {
		fmt.Println("Letter : ", string(letter))
	}
}
