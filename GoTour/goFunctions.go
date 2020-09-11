package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Time Start: [", time.Now(), "] ")
	fmt.Println("Hello, 世界")
	fmt.Println("Random Num Pre-Seed 01: ", rand.Intn(30))
	rand.Seed(16)
	fmt.Println("Random Num Post-Seed 01: ", rand.Intn(30))
	rand.Seed(16)
	fmt.Println("Random Num Post-Seed 02: ", rand.Intn(30))
	// fmt.Println("Random Num Post-Seed 02: ", rand.Intn(30))
	// fmt.Println("Random Num Post-Seed 02: ", rand.Intn(30))
	// fmt.Println("Random Num Post-Seed 02: ", rand.Intn(30))
	fmt.Println("Time End: [", time.Now(), "] ")

	x := 2
	y := 4
	fmt.Println(add1(x, y))
	fmt.Println(add2(x, y))

	// Test change
	fmt.Println()
	fmt.Println("Test change for git commit and push, again")

	/*
		Multiple Results:
		We can create a swap function that will allow us to control the
		order by which elements are returned from a function
	*/
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	z := 6
	c, d, e := swap2(x, y, z)
	fmt.Println(c, d, e)

	/*
		Named return values:
		Go's return values may be named. If they are named,
		they are treated as variables defined at the the top of the
		function..

		These names should be used to document the meaning of the
		return values..

		A return statement without arguments returns the named return
		values. This is known as a "naked return".

		Naked return statements should be used only in short functions,
		as with the example shown here.
		They can harm readability in longer functions
	*/

	// fmt.Println(split1(18))
	fmt.Print("split1 results: ")
	fmt.Println(split1(18))

	// alternatively, instantiate the variables ahead of time
	c, d = split2(18)
	fmt.Print("split2 results: ")
	fmt.Println(c, d)

	var1, var2, var3 := split3(18)
	fmt.Print("split3 results: ")
	fmt.Println(var1, var2, var3)

}

/*
Named return values function, we will create a split1(
) function to demonstrate
*/

func split1(sum int) (x, y int) {
	x = sum * 4 / 16
	y = sum - x

	return
}

func split2(sum int) (x, y int) {
	x = sum * 4 / 8
	y = sum - x

	return
}

func split3(intSumIn int) (x, y, z int) {
	x = (intSumIn * 2) / 8
	y = intSumIn - (x / 2)
	// y = (intSumIn * 4) % 4
	z = intSumIn - y
	// z = (intSumIn * 6) % 4

	return
}

/*
Multiple Results:
A function can return any number of results..
E.g. the 'swap' function below returns two strings
*/

// 1.
// '(string, string)' are the return types, unlike Java,
// placed after function parameters
func swap(x, y string) (string, string) {

	// Change order of element return
	// return x, y
	return y, x
}

// 2.
//
func swap2(x, y, z int) (int, int, int) {
	// return x, y, z
	// return x, z, y
	return y, x, z
}

// Addition Long version
func add1(x int, y int) int {
	return x + y
}

// Addition Shorter version
// '(int)' is the return type
func add2(x, y int) int {
	return x + y
}
