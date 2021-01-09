package main

import "fmt"

func main() {
	var point = 4

	switch point {
	case 8:
		fmt.Println("perfect")
	// Sebuah case dapat menampung banyak kondisi.
	case 7, 6, 5:
		fmt.Println("awesome")
	default:
		/*
			Tanda kurung kurawal ({ }) bisa diterapkan pada keyword case dan default
			Bagus jika dipakai pada blok kondisi yang didalamnya ada banyak statement,
			 kode akan terlihat lebih rapi dan mudah di-maintain.
		*/
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}
}
