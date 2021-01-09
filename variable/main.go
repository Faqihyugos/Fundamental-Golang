package main

import "fmt"

func main() {
	//var firstName string = "Faqih"

	//var lastName string
	//lastName = "Yugo Susilo"

	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	var firstName = "john"

	/* 	Tanpa var, tanpa tipe data, menggunakan perantara ":="
	Tanda := hanya digunakan sekali di awal pada saat deklarasi.
	Untuk assignment nilai selanjutnya harus menggunakan tanda =, contoh:
	*/
	lastName := "wick"
	lastName = "ethan"

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	// Multi Variabel
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	var fourth, fifth, sixth string = "empat", "lima", "enam"

	// type inference bisa berbeda tipe data nya
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	fmt.Println(first, second, third, fourth, fifth, sixth, seventh, eight, ninth)
	fmt.Println(one, isFriday, twoPointTwo, say)

	// Variable underscore
	_ = "belajar Golang"
	_ = "Golang itu mudah"
	name, _ := "john", "wick"
	fmt.Println(name)

}
