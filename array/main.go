package main

import "fmt"

func main() {
	// var languages [5]string
	// languages[0] = "Go"
	// languages[1] = "Ruby"
	// languages[2] = "Javascript"
	// languages[3] = "C#"
	// languages[4] = "Python"

	// fmt.Println(languages)
	// length := len(languages)
	// fmt.Println(length)

	languages := [...]string{
		"Ruby",
		"Python",
		"Go",
		"Javascript",
		"Kotlin",
		"C#",
	}

	for index, lang := range languages {
		fmt.Println("index : ", index, " language : ", lang)
	}
}
