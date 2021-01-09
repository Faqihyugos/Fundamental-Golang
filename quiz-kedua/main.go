package main

import "fmt"

func main() {
	title := "Golang the best language"

	// output index yang bernilai genap
	// for index, letter := range title {
	// 	if index%2 == 0 {
	// 		fmt.Println("Index : ", index, "Letter : ", string(letter))
	// 	}
	// }

	// output huruf vokal
	for index, letter := range title {
		letterString := string(letter)

		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("Index : ", index, "Letter : ", string(letter))
		}
	}
}
