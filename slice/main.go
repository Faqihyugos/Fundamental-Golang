package main

import "fmt"

func main() {
	gamingConsoles := []string{
		"PlayStation 4",
		"Nintendo switch",
		"Xbox One",
	}
	for _, console := range gamingConsoles {
		fmt.Println(console)
	}

}
