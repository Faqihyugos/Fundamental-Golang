package main

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("Halo, belajar golang")
	result := calculation.Add(8, 9)
	fmt.Printf("Hasil dari 8 + 9 = %d \n", result)
}
