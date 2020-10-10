package main

import (
	"fmt"
	"math"
)

type VrtxStrct01 struct {
	X, Y float64
}

// used with Abs03 below
type MyFloat64_01 float64

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

	// same function as above, but without the receiver argument
	fmt.Println()
	fmt.Println("Basic function - Abs02(v01): ", Abs02(v01))

	fmt.Println()
	fmt.Println("Simple get absolute value method (" +
		"ie a function with a receiver method): ")
	flt02 := MyFloat64_01(20)
	flt03 := math.Sqrt(81)
	fmt.Println("flt02.Abs03(-20): ", flt02.Abs03())
	fmt.Println("MyFloat64_01(flt03).Abs03(): ",
		MyFloat64_01(flt03).Abs03())
	fmt.Println("MyFloat64_01(-flt03).Abs03(): ",
		MyFloat64_01(-flt03).Abs03())

	// Methods with Pointer Receivers
	scaleVal := float64(10)

	fmt.Println()
	fmt.Println("ScaleNP method with non-Pointer Receiver")
	v02 := VrtxStrct01{
		3,
		4,
	}
	v02.ScaleNP(scaleVal)
	fmt.Println("v02.ScaleNP(scaleVal) ")
	fmt.Println("v02.Abs01(): ", v02.Abs01())

	fmt.Println()
	fmt.Println("ScaleWP method With Pointer Receiver")
	v03 := VrtxStrct01{
		3,
		4,
	}
	v03.ScaleWP(scaleVal)
	fmt.Println("v03.ScaleWP(scaleVal)")
	fmt.Println("v03.Abs01(): ", v03.Abs01())

	fmt.Println()
	fmt.Println("ScaleFunc function with Pointer parameter/argument" +
		" passed in")
	v04 := VrtxStrct01{
		3,
		4,
	}
	// original v04 values
	fmt.Println("v04.Abs01() before scale: ", v04.Abs01())
	// like other scale functions above,
	// this points to a vertex struct, and passes a scale value in,
	// to change the original value of the struct by the scale value
	ScaleFunc(&v04, scaleVal)
	// ScaleFunc(v04, scaleVal) // error if no pointer is passed
	fmt.Println("v04.ScaleWP(scaleVal)")
	fmt.Println("v04.Abs01() after scale: ", v04.Abs01())

}

/*

Methods

Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument,
(? ie like an incoming parameter profile?)

**The receiver appears in its own argument list between the func
keyword and the method name.

In this example, the Abs01 method has a receiver of type Vertex named v.

*/

// Declaring a method (ie a function with a receiver argument) on a
// struct type
// Function receiver argument: v1 VrtxStrct01
// 	(A receiver argument needs to be defined in the same scope as the
// method.)
// Function name: Abs()
// Function return data type: float64
// -This type of function, a method,
// puts the function 'on' the receiving data type's variable instance,
// see below v1.Abs01()

func (vrtx VrtxStrct01) Abs01() float64 {
	absVal := math.Sqrt(vrtx.X*vrtx.X + vrtx.Y*vrtx.Y)
	return absVal
}

// The same function without a receiver argument,
func Abs02(vIn VrtxStrct01) float64 {
	return math.Sqrt(vIn.X*vIn.X + vIn.Y*vIn.Y)
}

// Simple get absolute value method
// declaring method on non-struct type

func (fIn MyFloat64_01) Abs03() float64 {
	// This checks to see if fIn is negative, and if so,
	// returns the negative of that to get the absolute value
	if fIn < 0 {
		return float64(-fIn)
	}
	return float64(fIn)

}

// ScaleNP function with non-Pointer Receiver VrtxStrct01
// Will not change vrtx01's value??
func (vrtx VrtxStrct01) ScaleNP(f float64) {
	vrtx.X = vrtx.X * f
	vrtx.Y = vrtx.Y * f
}

// ScaleNP function with Pointer Receiver *VrtxStrct01
func (vrtx *VrtxStrct01) ScaleWP(f float64) {
	vrtx.X = vrtx.X * f
	vrtx.Y = vrtx.Y * f

}

// ScaleFunc is a normal function that receives a pointer receiver
// *VrtxStrct01, no receiver argument
func ScaleFunc(vrtx *VrtxStrct01, f float64) {
	vrtx.X = vrtx.X * f
	vrtx.Y = vrtx.Y * f
}
