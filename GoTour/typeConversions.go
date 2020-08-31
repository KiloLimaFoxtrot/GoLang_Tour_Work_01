package main

import (
	"fmt"
	"math"
)

func main() {

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))

	fmt.Println("Type Conversions / Casting")
	fmt.Println()

	fmt.Println("I. " +
		"\nx, y, z before type conversion: ")
	fmt.Println()

	fmt.Printf("variable: x \ntype: %T \nvalue: %v\n", x, x)
	fmt.Println()

	fmt.Printf("variable: y \ntype: %T \nvalue: %v\n", y, y)
	fmt.Println()

	fmt.Printf("variable: f \ntype: %T \nvalue: %v\n", f, f)
	fmt.Println()

	//Below, casting type float64 f to type uint,
	//with below syntax assigning a new variable z to that casted value
	var z uint = uint(f)
	//different output formatting
	fmt.Println("II. " +
		"\nx, y, z after type conversion: ")
	fmt.Println()
	
	fmt.Printf("variable: x, type: %T, value: %v\n", x, x)
	fmt.Printf("variable: y, type: %T, value: %v\n", y, y)
	fmt.Printf("\nvariable: z, type: %T, value: %v\n", z, z)
	fmt.Println()

	typeRemoveTest()

}

func typeRemoveTest() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	//with type conversion
	//fmt.Println()
	fmt.Println("III. " +
		"\nType Remove Test function output:")
	fmt.Println()

	fmt.Println("Type conversion present-")
	fmt.Println(x, y, z)
	fmt.Printf("variable: x, type: %T, value: %v\n", x, x)
	fmt.Printf("variable: y, type: %T, value: %v\n", y, y)
	fmt.Printf("variable: z, type: %T, value: %v\n", z, z)
	fmt.Println()

	//The below removes the type conversion of f to uint,
	//and stored as z, and just prints f
	//fmt.Println()
	fmt.Println("Type conversion absent-")
	fmt.Println(x, y, f)
	fmt.Printf("variable: x, type: %T, value: %v\n", x, x)
	fmt.Printf("variable: y, type: %T, value: %v\n", y, y)
	fmt.Printf("variable: f, type: %T, value: %v\n", f, f)
	fmt.Println()

}
