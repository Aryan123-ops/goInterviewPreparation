package main

import "fmt"

func main() {
	n := 7
	result := fibonacci(n)
	fmt.Println("Fibonacci of", n, "is", result)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
