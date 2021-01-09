package main

import (
	"fmt"
	"quiz-pertama/calculation"
)

func main() {
	fmt.Println("Hasil Soal Quiz Perkalian")
	result := calculation.Multiply(8, 5)
	fmt.Println(result)
}
