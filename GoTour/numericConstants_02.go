package main

import "fmt"

//Numeric constants are high precision values:
//An untyped constant takes the type needed by its context.

//Try printing needInt02(Big) too
//An int can be store at maximum a 64-bit integer,
//and sometimes less (sometimes ??)
//E.g.
//variable constant list declaration syntax
//BINARY to DECIMAL conversions:
// 1.
//Big1 = 1 << 100

//The below creates a large number by declaring a 1,
//and putting 100 zeros to the right of it,
//-i.e. by declaring a 1 and shifting it left 100 places away
//from zero.
//-i.e. the below assigns to the variable Big the value of
//the binary number that is 1 followed by 100 zeros.

// 2.
//Small1 = Big1 >> 99

//Shift the above 1 right 99 places,
//and we end up with 1 << 1,
//-ie 1 followed by 1 zero, the binary value 10,
//-ie the decimal value 2
//

const (
	//Try 01, seems too large
	// 1.
	Big1 = 1 << 100
	// 2.
	Small1 = Big1 >> 99

	//Try 02, smaller values by a degree of 10
	// 1.
	Big2 = 1 << 10
	// 2.
	Small2 = Big1 >> 9
)

func main() {
	fmt.Println()

	//To display original variable details,
	fmt.Println("Ok, Numeric Constants!" +
		"BINARY to DECIMAL conversions:")

	//but getting the following error...
	/*
		# command-line-arguments
		./numericConstants_02.go:58:59: constant 1267650600228229401496703205376 overflows int
		./numericConstants_02.go:59:12: constant 1267650600228229401496703205376 overflows int

		Compilation finished with exit code 2
	*/

	fmt.Println()

	/*
		fmt.Println("Original variable details: ")
		fmt.Printf("Variable Big1 type: %T, value: %v", needInt02(Big1),
			needInt02(Big1))
		fmt.Printf("Variable Small1 type: %T, value: %v", Small1, Small1)
		fmt.Println()*/

	//Try 01
	fmt.Println("needInt02: ")
	fmt.Println(needInt02(Small1))

	//The below won't work, as noted above,
	//An int can be store at maximum a 64-bit integer,
	//and sometimes less (sometimes ??)
	//fmt.Println(needInt02(Big1))
	fmt.Println()

	fmt.Println("needFloat02: ")
	fmt.Println(needFloat02(Small1))
	fmt.Println(needFloat02(Big1))
	fmt.Println()

	//Try 02
	/*fmt.Println("Original variable details: ")
	fmt.Printf("Variable Big2 type: %T, value: %v", Big2, Big2)
	fmt.Printf("Variable Small2 type: %T, value: %v", Small2, Small2)
	fmt.Println()

	fmt.Println("needInt02: ")
	fmt.Println(needInt02(Small2))
	fmt.Println(needInt02(Big2))
	fmt.Println()

	fmt.Println("needFloat02: ")
	fmt.Println(needFloat02(Small2))
	fmt.Println(needFloat02(Big2))
	fmt.Println()*/

}

//a function to retrieve the integer value of a numeric value passed

func needInt02(x int) int {
	return x*10 + 1
}

//a function to retrieve the float value of a numeric value passed
func needFloat02(x float64) float64 {
	return x * 0.1

}

//From the Tour of Go website
/*
func needInt02(x int) int {
	return x*10 + 1
}

func needFloat02(x float64) float64 {
	return x * 0.1
}
*/
