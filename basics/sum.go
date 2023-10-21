package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}

// func main() {
// 	num := 10
// 	sum := 0
// 	for i := 1; i <= num; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }

// func main() {
// 	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	sum := 0
// 	for i := 0; i < len(array); i++ {
// 		sum += array[i]
// 	}
// 	fmt.Println(sum)
// }

// func Sum(arr []int) int {
// 	sum := 0
// 	for i := 0; i < len(arr); i++ {
// 		sum += arr[i]
// 	}
// 	return sum
// }

// func main() {
// 	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	fmt.Println(Sum(array))
// }
