package main

import "fmt"

func main() {
	num := 1

	for {
		if num > 10 {
			break
		}
		fmt.Printf("%d\n", num)
		num++
	}
}
