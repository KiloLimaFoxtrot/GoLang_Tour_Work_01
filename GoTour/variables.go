package main

import (
	"fmt"
	"math/cmplx"
)

/*
*** FIELDS
 */

// Notes:
// I. Creating variables
// A. var 'instantiation' statement to declare a list of variables,
// or
// B. var'instantiation' individually ( like with Java or python)

// II. Can declare/instantiate variables at the package (
// ie like at the class field level of a java file) or at the
// function level,

// Examples:
// IA. var 'instantiation' statement, at package/field level
// uses a variable list
var (
	// instantiated, but non-initialized
	i               int  // default value is 0
	c, python, java bool // default value is false?

	// instantiated, and initialized
	// 1. Can omit the data type if we have an initializer value
	j, k, l, str1 = /*type*/ 1, 2, false, "str1 String!"

	// 2. Or can use the data type, eg bool, uint64 etc..
	ToBe bool = false
)

// IB.var 'instantiation' individually, at package/field level
// omiting data type
var typeOmmtd = 20

// using data type
var MaxInt uint64 = 1<<64 - 1
var z complex128 = cmplx.Sqrt(-5 + 12i)

// Default variables examples
// variables declared without an explicit initial value are
// given their zero/default value

//	numeric types: 0
//	boolean types: false
//	String types: "" //ie the empty string
var (
	int1 int
	flt1 float64
	bol1 bool
	str2 string
)

/*
*** MAIN SECTION
 */
func main() {
	fmt.Println()
	fmt.Println("Variables Go!: ")
	basicTypes01()
	basicTypes02()
	zeroValues01(int1, flt1, bol1, str2)

}

/*
*** FUNCTIONS
 */

// declaring variable types in parameter field of function
func zeroValues01(theInt1 int, theFlt1 float64, theBol1 bool,
	theStr2 string) {

	fmt.Println("Variables' default values single print statement" +
		" 01:")
	fmt.Println(theInt1, theFlt1, theBol1, theStr2)

	fmt.Println("\nVariables' default values individual print" +
		" statements 01:")
	fmt.Printf("Default value for %T = %v", theInt1, theInt1)
	fmt.Printf("\nDefault value for %T = %v", theFlt1, theFlt1)
	fmt.Printf("\nDefault value for %T = %v", theBol1, theBol1)
	// %q ?? used below
	fmt.Printf("\nDefault value for %T = %q\n", theStr2, theStr2)

	// or
	fmt.Println("\nVariables' default values single print" +
		" statement 02:")
	fmt.Printf("%v %v %v %q\n", theInt1, theFlt1, theBol1, theStr2)

}

func basicTypes02() {
	// basic variable data types
	/*
		bool
		string
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint16, uintptr

		byte // alias for uint8

		rune // alias for int32
			// represents a unicode code point

		float32, float64

		complex64, complex128

		The following example shows variables of several types,
		and also that variable declarations may be "factored" into
		blocks, as with import statements

		The int, uint,and uintptr types are usually 32 bits wide on a
		32-bit system and 64 bits wide on 64-bit systems.

		When you need an integer value you should use int unless you
		have a specific reason so use a sized or unsigned integer type.

		//See above in field section/area of this 'class'
		var (
		ToBe bool = false
		MaxInt uint64 = 1<<64-1
		z complex128 = cmplx.Sqrt(-5 + 12i)
		)

		below we use string formatting to print a variable's data
		type, %T calls that, and the variables value, %v calls that.
	*/

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)

	fmt.Printf("Type: %T Value: %v\n", typeOmmtd, typeOmmtd)

	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)

	fmt.Printf("Type: %T Value: %v\n", z, z)

}

func basicTypes01() {

	// var statement to declare a list of variables at function level,
	// var i int //default value is 0

	// Printing the variables
	fmt.Println(i, c, python, java)

	// variables and initializers
	// If an initializer is present,
	// we can omit the variable data type such as integer because
	// the GoLang will use the type of the initializer

	// var j, k, l, str1 = 1, 2, false, "str1 String!"
	fmt.Println(j, k, l, str1)

	// Short variable declarations
	// Inside a function,
	// the ':=' short assignment statement can be used in place of a
	// var declaration with implicit type

	// Outside a function, every statement begins with a keyword (
	// var, func, and so on) and so the ':=' construct is not possible

	var i2, j2 int = 1, 2
	k2 := 3
	c2, python2, java2 := true, false, "no!"

	fmt.Println(i2, j2, k2, c2, python2, java2)

}

/*
*** ToString ??
 */
