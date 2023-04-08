package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		sum := 0
		for i := 0; i < 100; i++ {
			sum += i
		}
		c <- sum
	}()

	result := <-c
	fmt.Println(result)
}

// In this example, we create a channel c and then launch a goroutine to perform a computation.
// The computation calculates the sum of the numbers from 0 to 999,999, and then sends the result
// over the channel c. The main goroutine then waits to receive the result from c and prints it
// out.

// This simple example demonstrates how goroutines and channels can be used to perform concurrent
// computations without requiring explicit locking or synchronization mechanisms. Goroutines can
//  communicate via channels without the need for mutexes, condition variables, or other
// traditional
// concurrency primitives. This makes it easier to write correct concurrent code and to take
// advantage of modern hardware architectures.
