package main

import "fmt"

func main() {
	// var myMap map[string]int
	// myMap = map[string]int{}

	// myMap["Ruby"] = 9
	// myMap["Go"] = 9
	// myMap["JavaScript"] = 8

	// fmt.Println(myMap)

	myMap := map[string]string{"ruby": "is beautiful", "go": "is super fast"}

	for index, Map := range myMap {
		fmt.Println(index, ":", Map)
	}

}
