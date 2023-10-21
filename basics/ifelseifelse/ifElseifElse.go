package main

import "fmt"

func main() {
	num := 99

	if num < 50 {
		fmt.Println(num, "is smaller than 50")
		return
	} else if num >= 51 && num <= 98 {
		fmt.Println(num, "is in between 50 to 100")
		return
	} else {
		fmt.Println("none of the condition matches")
	}
}
