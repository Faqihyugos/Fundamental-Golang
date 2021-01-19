package main

import "fmt"

func main() {
	students := []map[string]string{
		{"name": "Faqih", "score": "A"},
		{"name": "Yugo", "score": "B"},
		{"name": "Mario", "score": "E"},
	}

	for _, student := range students {
		fmt.Println(student["name"], "-", student["score"])
	}
}
