package main

import "fmt"

func main() {
	mul := 1

	for mul <= 10 {
		product := 5 * mul

		fmt.Printf("5 * %d = %d\n", mul, product)
		mul++
	}
}
