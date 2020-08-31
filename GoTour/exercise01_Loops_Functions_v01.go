package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("")
//	fmt.Println("Exercise 01: Loops and Functions. ")
//
//	/*
//		As a way to play with functions and loops,
//		lets implement a square root function:
//
//		Square Root Function Problem Description:
//		-Computers typically compute the square root of x3 by using a
//		loop.
//		-Starting with some guess z1,
//		we can adjust z1 based on how close z1^2 is to x3,
//		producing a better guess
//
//		z1 -= (z1*z1 - x3) / (2*z1)
//
//		Breakdown of above equation, known as "Newton's Method":
//		-The (x3 * x3 - z1) ie x3^2 - z1,
//		is how far away x3^2 is from where it needs to be (x3)
//		-The division by (2*z1) is the derivative of z1^2,
//		which is to scale how much we adjust z1 by how quickly z1^2 is
//		changing..
//		-This method works well for many functions,
//		especially square root
//
//	*/
//
//	//number to find square root of
//	var x1 float64 = 81
//	//square root value of x3
//	var z1 float64 = 1
//
//	//## squareRoot01 function
//	//function to find the square root z1 of x3
//	fmt.Println()
//	var n1 float64 = squareRoot01(x1, z1)
//	fmt.Printf("The square root of %v is %v. \n", x1, n1)
//
//	//## squareRoot02 function
//	fmt.Println()
//	var n2 float64 = squareRoot02(x1, z1)
//	fmt.Printf("The square root of %v is %v. \n", x1, n2)
//
//	//##** Sqrt function
//	fmt.Println()
//	var n3 float64 = Sqrt(x1, z1)
//	fmt.Printf("The square root of %v is %v. \n", x1, n3)
//
//	//End spacer
//	fmt.Println()
//}
//
//func Sqrt(theX float64, theZ float64) float64 {
//
//	var zTemp = theZ
//
//	for i := theZ; zTemp-i >= 0; i++ {
//		//Original equation from example..
//		zTemp -= (zTemp*zTemp - theX) / (2 * zTemp)
//		fmt.Printf("zTemp val: %v,  ", zTemp)
//		fmt.Printf("i val: %v\n", i)
//	}
//
//	return zTemp
//}
//
//func squareRoot02(theX float64, theZ float64) float64 {
//
//	var zTemp = theZ
//
//	for i := theZ; i < theZ+10 && zTemp-i >= 0; i++ {
//		//Original equation from example..
//		zTemp -= (zTemp*zTemp - theX) / (2 * zTemp)
//		fmt.Printf("zTemp val: %v,  ", zTemp)
//		fmt.Printf("i val: %v\n", i)
//	}
//
//	return zTemp
//}
//
//func squareRoot01(theX float64, theZ float64) float64 {
//
//	var zTemp float64
//
//	for i := theZ; i < theZ+10; i++ {
//		zTemp = i
//		//Original equation from example..
//		zTemp -= (zTemp*zTemp - theX) / (2 * zTemp)
//		fmt.Printf("zTemp val: %v, ", zTemp)
//		fmt.Printf("i val: %v\n", i)
//	}
//
//	return zTemp
//}
