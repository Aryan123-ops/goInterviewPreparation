package main

import "fmt"

// func main() {
// 	n := 5
// 	for i := 0; i < n; i++ {
// 		for j := 0; j <= i; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

func main() {
	i := 0
	for i <= 10 {
		fmt.Printf("%d ", i)
		i += 2
	}
}
