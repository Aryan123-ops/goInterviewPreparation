package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := []string{"hi", "How", "are", "you"}
	// str := "Welcome"
	// str2 := "katta lovers"

	// res1 := strings.Contains(str1, "Weo=")
	// res2 := strings.Contains(str2, "vers")

	result := IfArrayContains(str1)
	// result := ifStringContains(str)
	fmt.Println(result)
	// fmt.Println(res2)
}

func IfArrayContains(s []string) bool {
	for _, value := range s {
		if strings.Contains(value, "hi") {
			return true
		} else {
			return false
		}
	}
	return false
}

// func ifStringContains(s string) bool {
// 	res := strings.Contains(s, "Wel")
// 	return res
// }
