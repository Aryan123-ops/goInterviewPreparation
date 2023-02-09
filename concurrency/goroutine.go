package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()
	time.Sleep(5 * time.Second)
	go byebye()
	time.Sleep(5 * time.Second)
	fmt.Println("okay bye")
}

func hello() {
	fmt.Println("hello ")
}

func byebye() {
	fmt.Println("Bye Bye")
}
