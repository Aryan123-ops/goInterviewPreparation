package main

import (
	"fmt"
)

// Create a struct and provide fileds that will be used in Method
type Triangle struct {
	Height float32
	Base   float32
}

// create an interface that will store some method signatures.
type MyInterface interface {
	Area() float32
}

// Define a method that will be used in Interface.
func (t Triangle) Area() float32 {
	return 0.5 * t.Base * t.Height
}

// Call everything in main()
func main() {
	// initialised variable of type struct and provide the value to the struct fields.
	triangleObject := Triangle{Base: 10, Height: 10}

	// create a variable of type interface.
	var shapeObject MyInterface

	// compare interface variable and struct variable
	shapeObject = triangleObject

	// prints the outout by calling interfaceVariable.MethodName()
	fmt.Println("Triangle Area", shapeObject.Area())
}
