package main

import "fmt"

var Data interface{} = 42

func main() {
	if value, ok := Data.(int); ok {
		fmt.Println("Dat is an int", value)
		return
	} else if value, ok := Data.(string); ok {
		fmt.Println("Data is string", value)
		return
	} else if value, ok := Data.(float64); ok {
		fmt.Println("data is float", value)
		return
	} else {
		fmt.Println("Data is none of the int,string, float64")
		return
	}

}
