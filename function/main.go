package main

import "fmt"

func main() {
	// sentence := printMyResult("Saya sedang")
	// fmt.Println(sentence)
	result := add(20, 10)
	fmt.Println(result)
}

func add(number, numberTwo int) int {
	return number + numberTwo
}

// func printMyResult(sentence string) string {
// 	newSentence := sentence + " Belajar Golang"
// 	return newSentence
// }

// 1. input
// 2. proses
// 3. output
