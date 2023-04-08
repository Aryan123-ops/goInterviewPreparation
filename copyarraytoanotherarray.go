// Go program to illustrate how
// to copy an array by value
package main

import "fmt"

func main() {
	my_arr1 := [5]string{"Scala", "Perl", "Java", "Python", "Go"}
	my_arr2 := my_arr1

	fmt.Println("Array_1: ", my_arr1)
	fmt.Println("Array_2:", my_arr2)

}
