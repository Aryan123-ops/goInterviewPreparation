package main

import "fmt"

func main() {
	num := 10

	if num%2 == 0 {
		fmt.Println(num, "is Even")
		return // it just returns the result and the program will immediately exists the main
		// function and terminates.
	} else {
		fmt.Println(num, "is odd")
	}
}
