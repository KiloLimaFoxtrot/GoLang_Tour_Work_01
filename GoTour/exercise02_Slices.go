package main

import (
	"fmt"
	// "golang.org/x/tour/v0.0.0-20200508155540-0608babe047d"
)

func main() {
	
	fmt.Println()
	fmt.Println("Exercise 02: Slices")
	
	// Implement Pic. It should return a slice of length dy,
	// each element of which is a slice of dx 8-bit unsigned integers.
	// When you run the program, it will display your picture,
	// interpreting the integers as grayscale (well, bluescale) values.
	//
	// The choice of image is up to you.
	// Interesting functions include (x+y)/2, x*y, and x^y.
	//
	// (You need to use a loop to allocate each []uint8 inside the
	// [][]uint8.)
	//
	// (Use uint8(intValue) to convert between types.)
	
	// Pic.Show(Pic) //??
	
	// lenX := 10
	// lenY := 10
	// picArry := Pic(dx, dy, lenX, lenY)
	dx := 2
	dy := 2
	picArry := Pic(dx, dy)
	
	fmt.Println()
	print2DArry(picArry)
	
}

// func Pic(dx, dy, lenX, lenY int) [][]uint8 {
func Pic(dx, dy int) [][]uint8 {
	x := dx
	y := dy
	lenX := 10
	lenY := 10
	
	picArryOut := make([][]uint8, lenX)
	// Syntax reference in notes with with a bullet of
	// 'exercise02_Slices.go'
	for i := 0; i < 10; i++ {
		picArryOut[i] = make([]uint8, lenY)
		for j := 0; j < 10; j++ {
			val := uint8(x * y)
			picArryOut[i][j] = val
			y++
		}
		x++
	}
	return picArryOut
}

/*
// GoTour site work for the above problem
//
//
// package main
//
// import "golang.org/x/tour/pic"
//
// func main() {
// 	pic.Show(Pic)
// }
//
// func Pic(dx, dy int) [][]uint8 {
// 		picArryOut := make([][]uint8, dy)
// 	for i := 0; i < dy; i++ {
// 		picArryOut[i] = make([]uint8, dx)
// 		for j := 0; j < dx; j++ {
// 			val := uint8(i * j)
// 			picArryOut[i][j] = val
// 		}
// 	}
// 	return picArryOut
// }
*/

/*
Solution from github:
(https://gist.github.com/fonglh/d5e556ad14cf36f7c6a5)

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	myPic := make([][]uint8, dy)
	//allocate
	for i := range myPic {
		myPic[i] = make([]uint8, dx)
	}

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			myPic[x][y] = uint8((x+y)/2)
		}
	}
	return myPic
}

func main() {
	pic.Show(Pic)
}

*/


/*
To format a string we need the %s formatter.
fmt.Printf("%s", "a string")    // a string

// format the size of string
fmt.Printf("|%8s|", "apple")   // |   apple|

*/

func print2DArry(Arry2D [][]uint8) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("[%3v]", Arry2D[i][j])
		}
		fmt.Println()
	}
	
}
