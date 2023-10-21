package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["avinash"] = 66
	m["vinash"] = 67
	m["Abhilash"] = 68
	m["Prakash"] = 69

	for k, v := range m {
		fmt.Println("key is :", k+"   value is :", v)
	}
	// fmt.Println(m["Prakash"])
}
