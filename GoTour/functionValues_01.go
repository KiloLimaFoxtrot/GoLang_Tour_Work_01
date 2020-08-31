package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println()
	fmt.Println("****Function Values:")
	
	// Functions are values too.
	// They can be passed around just like other values.
	//
	// Function values may be used as
	// 1. function arguments and/or
	// 2. return values.
	
	// ? In the former 1. option,
	// does that mean it is possible to pass the arguments of a first
	// function to second function?
	//
	// How does one pass values to that function,
	// which would be in addition to passing values?
	
	// The values are passed to the second function in one way by
	// using the return statement of the second function pass values
	// to a call of the second function
	
	// this is a function within a function, ie within main function,
	// and sets the function call return value to a variable 'hypot01'
	
	// note: the below internal function does not have a name,
	// it just has arguments to perform an operation,
	// which specifically are to take two float64 values from x and y,
	// and return a float64 value from calculation of the hypotenuse
	// math.Sqrt(x*x + y*y)
	hypot01 := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y) // returns the hypotenuse
	}
	
	fmt.Println()
	// to call the internal function that is within this main
	// function and pass simple values
	fmt.Println("hypot01(5, 12): ", hypot01(5, 12))
	
	// Calls to the external functions...
	
	// This (1.) call to compute02 passes the arguments of this
	// internal function to compute02: 'compute02(hypot01)'
	// 1.
	fmt.Println("compute02(hypot01): ", compute01(hypot01))
	
	// This call passes the arguments
	// 2.
	fmt.Println("compute02(math.Pow): ", compute01(math.Pow))
	
}

func compute01(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
