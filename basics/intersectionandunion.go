package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7}
	fmt.Println("intersection of two arrays are:", intersection(slice1, slice2))
	fmt.Println("union of two arrays are:", union(slice1, slice2))

}

func intersection(slice1, slice2 []int) (intersec []int) {
	mapp := make(map[int]bool)
	for _, item := range slice1 {
		mapp[item] = true
	}
	for _, item := range slice2 {
		if _, ok := mapp[item]; ok {
			intersec = append(intersec, item)
		}
	}
	return
}

func union(slice1, slice2 []int) []int {
	mapp := make(map[int]bool)

	for _, item := range slice1 {
		mapp[item] = true
	}
	for _, item := range slice2 {
		if _, ok := mapp[item]; !ok {
			slice1 = append(slice1, item)
		}
	}
	return slice1
}
