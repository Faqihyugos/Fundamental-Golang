package main

import "fmt"

func main() {
	// sentence := printMyResult("Saya sedang")
	// fmt.Println(sentence)

	// result := add(20, 10)
	// fmt.Println(result)

	luas, keliling := calculate(10, 2)
	fmt.Println("Luas : ", luas)
	fmt.Println("Keliling : ", keliling)

}

func calculate(panjang, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}

// func add(number, numberTwo int) int {
// 	return number + numberTwo
// }

// func printMyResult(sentence string) string {
// 	newSentence := sentence + " Belajar Golang"
// 	return newSentence
// }

// 1. input
// 2. proses
// 3. output
