package main

import "fmt"

func main() {
	fmt.Println("")
	fmt.Println("Exercise 01: Loops and Functions. ")

	//number to find square root of
	var x1 float64 = 81
	//square root value of x3
	var z1 float64 = 1

	//##** Sqrt function
	fmt.Println()
	var n3 float64 = Sqrt(x1, z1)
	fmt.Printf("The square root of %v is %v. \n", x1, n3)

	//End spacer
	fmt.Println()
}

func Sqrt(theX float64, theZ float64) float64 {

	var zTemp = theZ

	for i := theZ; zTemp-i >= 0; i++ {
		//Original equation from example..
		zTemp -= (zTemp*zTemp - theX) / (2 * zTemp)
		fmt.Printf("zTemp val: %v,  ", zTemp)
		fmt.Printf("i val: %v\n", i)
	}

	return zTemp
}
