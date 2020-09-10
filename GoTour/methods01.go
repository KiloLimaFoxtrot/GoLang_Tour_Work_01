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

// Declaring a method (ie a function with a receiver argument) on a
// struct type
// Function reciever argument: v1 VrtxStrct01
// Function name: Abs()
// Function return data type: float64
// -This type of function, a method,
// puts the function 'on' the receiving data type's variable instance,
// see below v1.Abs01()
func (vIn VrtxStrct01) Abs01() float64 {
	absVal := math.Sqrt(vIn.X*vIn.X + vIn.Y*vIn.Y)
	return absVal
}

// The same function without a receiver argument,
func Abs02(vIn VrtxStrct01) float64 {
	return math.Sqrt(vIn.X*vIn.X + vIn.Y*vIn.Y)
}

// Simple get absolute value method
// declaring method on non-struct type
type MyFloat01 float64

func (fIn MyFloat01) Abs03() float64 {
	// This checks to see if fIn is negative, and if so,
	// returns the negative of that to get the absolute value
	if fIn < 0 {
		return float64(-fIn)
	}
}

func main() {
	fmt.Println()
	fmt.Println("**** Methods 01:")
	
	v01 := VrtxStrct01{
		X: 3,
		Y: 4,
	}
	
	// Here is a function Abs01(
	// ) 'on' the data type VrtxStrct01 variable instance v01
	v1AbsVal := v01.Abs01()
	fmt.Println()
	fmt.Println("Method on variable assigned to v1AbsVal: ", v1AbsVal)
	
	// condensed version of just above,
	// removes the variable initialization
	fmt.Println()
	fmt.Println("Method on variable - v01.Abs01(): ", v01.Abs01())
	
	fmt.Println()
	fmt.Println("Basic function - Abs02(v01): ", Abs02(v01))
	
}
