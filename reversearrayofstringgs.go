package main

import "fmt"

func main() {
	strings := []string{"foo", "bar", "baz", "qux"}
	fmt.Println("Original array:", strings)

	for i := 0; i < len(strings)/2; i++ {
		j := len(strings) - i - 1
		strings[i], strings[j] = strings[j], strings[i]
	}

	fmt.Println("Reversed array:", strings)
}
