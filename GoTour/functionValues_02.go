package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println()
	fmt.Println("****Function Values 02:")

	// Stp01: make an internal function

	// using the GoTour's example (
	// no comma after last variable in parameter signature,
	// I always forget that : )
	hypot02 := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	// initializing variables
	var (
		x01, x02, y01, y02 float64
	)

	fmt.Println()
	fmt.Println("hypot02(x01, y01): ")
	x01 = 5
	y01 = 12
	hypot02Call01 := hypot02(x01, y01)
	fmt.Println("hypot02(x01, y01): ", hypot02Call01)

	fmt.Println()
	fmt.Println("compute02(hypot02): ")
	compute02(hypot02)

	fmt.Println()
	fmt.Println("compute02(math.Pow): ")
	compute02(math.Pow)

	fmt.Println()
	fmt.Println("compute03(math.Pow, x02, y02): ")
	x02 = 3
	y02 = 4
	compute03(math.Pow, x02, y02)

	fmt.Println()
	compute04Call01 := compute04(hypot02, x01, y01)
	fmt.Println("func compute04, hypot02, fn04(x01, y01): ",
		compute04Call01)
	fmt.Println("x01: ", x01)
	fmt.Println("y01: ", y01)

	fmt.Println()
	compute04Call02 := compute04(hypot02, x02, y02)
	fmt.Println("func compute04, hypot02, fn04(x02, y02): ",
		compute04Call02)
	fmt.Println("x02: ", x02)
	fmt.Println("y02: ", y02)

	fmt.Println()
	compute04Call03 := compute04(math.Pow, x02, y02)
	fmt.Println("func compute04, math.Pow, fn04(x02, y02): ",
		compute04Call03)
	fmt.Println("x02: ", x02)
	fmt.Println("y02: ", y02)

}

func compute02(fn02 func(float64, float64) float64) {
	fmt.Println("func compute02 fn02(3, 4): ", fn02(3, 4))
}

// Function (here, 'compute03') that receives both
// 1. from another function (here, fn03) the arguments
// (here, func(float64, float64) float64)
// and values to plug into the arguments
// 2. (here, x float64, y float64)
func compute03(fn03 func(float64, float64) float64, x float64,
	y float64) {
	fmt.Println("func compute03 fn03(x, y): ", fn03(x, y))
	fmt.Println("x00: ", x)
	fmt.Println("y00: ", y)
}

func compute04(fn04 func(float64, float64) float64, x float64,
	y float64) float64 {

	return fn04(x, y)

}
