package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Take two strings
	str1 := "listen"
	str2 := "silent"

	// Replace space of given strings
	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")

	//split strings to slice of character
	slice1 := strings.Split(str1, "")
	slice2 := strings.Split(str2, "")

	//Sort slice of character
	sort.Strings(slice1)
	sort.Strings(slice2)

	//check if sorted slice of character is equal
	if strings.Join(slice1, "") == strings.Join(slice2, "") {
		fmt.Println("Anagram")
	} else {
		fmt.Println("Not Ana")
	}
}
