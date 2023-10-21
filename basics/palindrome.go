package main

import (
	"fmt"
)

func main() {
	str := "racecar"

	if isPalindrome(str) {
		fmt.Println(str, "is palindrome")
	} else {
		fmt.Println(str, "not palindrome")
	}
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
