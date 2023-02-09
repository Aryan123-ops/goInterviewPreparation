package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{8, 4, 7, 6, 5, 3, 2, 1}
	sort.Ints(arr)
	fmt.Println("sorted in ascending order:", arr)
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println("sorted in descending order:", arr)
}
