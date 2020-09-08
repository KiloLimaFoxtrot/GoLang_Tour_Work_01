package main

import (
	"fmt"
	"math"
)

/*

Methods

Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument,
(? ie like an incoming parameter profile?)

The receiver appears in its own argument list between the func
keyword and the method name.

In this example, the Abs01 method has a receiver of type Vertex named v.

*/

type VrtxStrct01 struct {
	X, Y float64
}

// Function reciever argument: v1 VrtxStrct01
// Function name: Abs()
// Function return data type: float64
// -This type of function, a method,
// puts the function 'on' the receiving data type's variable instance,
// see below v1.Abs01()
func (v1 VrtxStrct01) Abs01() float64 {
	absVal := math.Sqrt(v1.X*v1.X + v1.Y*v1.Y)
	return absVal
}

// The same method without a receiver argument,
// written as a basic function

func Abs02(v1 VrtxStrct01) float64 {
	return math.Sqrt(v1.X*v1.X + v1.Y*v1.Y)
}

func main() {
	fmt.Println()
	fmt.Println("**** Methods 01:")

	v1 := VrtxStrct01{
		X: 3,
		Y: 4,
	}

	// Here is a function Abs01(
	// ) 'on' the data type VrtxStrct01 variable instance v1
	v1AbsVal := v1.Abs01()

	fmt.Println()
	fmt.Println("v1AbsVal: ", v1AbsVal)

}
