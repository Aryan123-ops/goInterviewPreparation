package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Enter a number:")
	_, err := fmt.Scanln(&num)

	if err != nil {
		fmt.Println(err)
	}

	if num%2 == 0 {
		fmt.Println(num, "is even")
		return
	} else {
		fmt.Println(num, "is odd")
	}
}
