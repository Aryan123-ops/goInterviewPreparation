package main

import (
	"fmt"
	"sync"
)

func first(wg *sync.WaitGroup) {
	defer wg.Done() // This decreases counter by 1
	fmt.Println("I am first runner")

}

func second(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("I am second runner")
}

func execute() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go first(wg)
	go second(wg)

	wg.Wait()
}

func main() {
	// Launching both the runners
	execute()
}
