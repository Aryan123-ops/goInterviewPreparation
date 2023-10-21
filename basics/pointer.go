package main

import "fmt"

func main() {
	x := 10         // declare and initialize a value
	y := &x         // declare a pointer that points to x's memory address
	fmt.Println(x)  // prints 10
	fmt.Println(y)  // prints the memory address of x
	fmt.Println(*y) // prints the value stored at the memory address that y points to (which is 10)
}
