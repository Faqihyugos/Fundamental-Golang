package main

import "fmt"

func main() {
	// var myMap map[string]int
	// myMap = map[string]int{}

	// myMap["Ruby"] = 9
	// myMap["Go"] = 9
	// myMap["JavaScript"] = 8

	// fmt.Println(myMap)

	// For Map
	// myMap := map[string]string{
	// 	"ruby":       "is beautiful",
	// 	"go":         "is super fast",
	// 	"JavaScript": "hype",
	// }

	// for key, value := range myMap {
	// 	fmt.Println("Key : ", key, ", Value : ", value)
	// }

	// fmt.Println("===============")

	// delete(myMap, "go")

	// for key, value := range myMap {
	// 	fmt.Println("Key : ", key, ", Value : ", value)
	// }

	// Melihat value map
	myMap := map[string]string{
		"ruby":       "is beautiful",
		"go":         "is super fast",
		"JavaScript": "hype",
	}

	value, isAvailabel := myMap["python"]

	fmt.Println(value)
	fmt.Println(isAvailabel)
}
