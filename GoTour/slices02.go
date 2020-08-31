package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	fmt.Println("Slices cntd...")
	fmt.Println("****Slice Defaults:")
	// When slicing,
	// you may omit the high or low bounds to use their defaults instead
	
	// The default is zero for the lower bound,
	// and the length of the of the slice for the high bound
	// e.g. for the array
	
	//  var arryX [10]int
	
	// these slice expressions are equivalent:
	// arryX[0:10]
	// arryX[:10]
	// arryX[0:]
	// arryX[:]
	
	fmt.Println()
	fmt.Println("  **FUNCTION slice02: ")
	slice02()
	
	fmt.Println()
	fmt.Println("  **FUNCTION slice03: ")
	slice03()
	
	fmt.Println()
	fmt.Println("****Slice length and capacity:")
	// Slice length and Capacity
	// A slice has both a length and a capacity.
	//
	// The length of a slice is the number of elements it contains.
	//
	// The capacity of a slice is the number of elements in the
	// underlying array, counting from the first element in the slice.
	//
	// The length and capacity of a slice s can be obtained using the
	// expressions len(s) and cap(s).
	//
	// You can extend a slice's length by re-slicing it,
	// provided it has sufficient capacity.
	//
	// Try changing one of the slice operations in the example program
	// to extend it beyond its capacity and see what happens.
	
	fmt.Println()
	fmt.Println("  **FUNCTION slice04: ")
	slice04()
	
	fmt.Println()
	fmt.Println("****Nil Slices:")
	// The zero value of a slice is nil.
	//
	// A nil slice has a length and capacity of 0 and has no
	// underlying array.
	
	fmt.Println()
	fmt.Println("  **FUNCTION slice05: ")
	slice05()
	
	fmt.Println()
	fmt.Println("****Creating a slice with make. ")
	
	// ** How to instantiate a dynamically sized array / slice of
	// specific length/and capacity, use make([]type, int length,
	// int capacity)
	
	// Slices can be created with the built-in make function; which is
	// how you create dynamically sized arrays
	
	// The make function allocates a zeroed array and returns a slice
	// that refers to that array
	
	// a := make([]int, 5) 		//here, len(a) = 5,
	// ie the length of (a)=5
	
	// To specify a capacity, pass a third argument to make:
	
	// b := make([]int, 0, 5) 	//here, len(b) = 0, cap(b) = 5
	
	// and/or with the below
	// b = b[:cap(b)] 			//here, len(b) = 5, cap(b) = 5
	// b = b[1:] 				//here, len(b) = 4, cap(b) = 4
	
	fmt.Println()
	fmt.Println("  **FUNCTION slice06: ")
	slice06()
	
	fmt.Println()
	fmt.Println("****Slices of slices:")
	
	// Slices can contain any type, including other slices
	
	// **** Working with 2D Arrays/Slices and printing/outputting via
	// iterating through the 2D Array/Slice
	
	fmt.Println()
	fmt.Println("  **FUNCTION slice07: ")
	slice07()
	
	fmt.Println()
	fmt.Println("****Appending to a slice:")
	
	// It is common to append new elements to a slice,
	// and so Go provides a built-in append function.
	// The documentation of the built-in package describes append.
	//
	// func append(s []T, vs ...T) []T
	//
	// The first parameter s of append is a slice of type T,
	// s []T -> is a slice of type T
	// and the rest are T values to append to the slice.
	//
	// The resulting value of append is a slice containing all the
	// elements of the original slice plus the provided values.
	//
	// If the backing array of s is too small to fit all the given
	// values a bigger array will be allocated.
	//
	// The returned slice will point to the newly allocated array.
	
	fmt.Println()
	fmt.Println("  **FUNCTION slice08: ")
	slice08()
	
}

// Helper Functions
func printSlice(arryNme string, arrySx []int) {
	fmt.Printf("%v: %v \n", arryNme, arrySx)
	fmt.Printf("%v length: %v, capacity: %v, \n", arryNme, len(arrySx),
		cap(arrySx))
}

func itratBildSlceArry(slceArry [][]string) {
	// Print the board, after each turn:
	// join elements?? and print a string form out
	for i := 0; i < len(slceArry); i++ {
		fmt.Printf("%s\n", strings.Join(slceArry[i], " "))
	}
}

// Array/Slice study functions

func slice08() {
	// make a slice of type int
	var slce01 []int
	fmt.Println()
	printSlice("slce01", slce01)
	
	// the append function will work on nil slices
	slce01 = append(slce01, 0)
	fmt.Println()
	printSlice("slce01", slce01)
	
	// here we are adding to a nil slice,
	// hence there is not sufficient capacity to take an element,
	// but in GoLang, in this case a larger Array will be allocated,
	// and the returned slice will point to the newly allocated larger
	// Array
	slce01 = append(slce01, 1)
	fmt.Println()
	printSlice("slce01", slce01)
	
	// We can add more than one element at a time,
	// even iterate through and add
	slce01 = append(slce01, 2, 3, 4)
	fmt.Println()
	printSlice("slce01", slce01)
	
	fmt.Println()
	var cnt int = slce01[len(slce01)-1]
	fmt.Println("slce01[len(slce01)-1]: ", slce01[len(slce01)-1])
	fmt.Printf("cnt %v \n", cnt)
	
	fmt.Println()
	var indx int = cnt + 1
	for i := 0; i < 10-indx; i++ {
		slce01 = append(slce01, cnt+1)
		fmt.Println(cnt + 1)
		printSlice("slce01", slce01)
		fmt.Println()
		cnt++
	}
	// printSlice("slce01", slce01)
	
	// for i := 0; i < 10-cap(slce01); i++ {
	//	slce01 = append(slce01, cnt+1)
	//	printSlice("slce01", slce01)
	//	cnt++
	// }
	// printSlice("slce01", slce01)
	
}

func slice07() {
	fmt.Println()
	fmt.Println("Create a tic-tac-toe board")
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	
	fmt.Println()
	fmt.Println("The players take turns")
	board[0][0] = "X"
	// Print the board, after each turn:
	// print a string form out, and join
	fmt.Println()
	fmt.Println("Round 01: X move")
	itratBildSlceArry(board)
	
	board[2][2] = "O"
	fmt.Println()
	fmt.Println("Round 02: O move")
	itratBildSlceArry(board)
	
	board[1][2] = "X"
	fmt.Println()
	fmt.Println("Round 03: X move")
	itratBildSlceArry(board)
	
	board[1][0] = "O"
	fmt.Println()
	fmt.Println("Round 04: O move")
	itratBildSlceArry(board)
	
	board[0][2] = "X"
	fmt.Println()
	fmt.Println("Round 05: X move")
	itratBildSlceArry(board)
	
}

func slice06() {
	
	// ** How to instantiate a dynamically sized array / slice of
	// specific length/and capacity, use make([]type, int length,
	// int capacity)
	
	arryA := make([]int, 5) // here, a length = 5
	fmt.Println()
	printSlice("arryA", arryA)
	
	arryB := make([]int, 0, 5) // here, arryB length = 0, capacity = 5
	fmt.Println()
	printSlice("arryB", arryB)
	
	arryC := arryB[:2] // here, only indexes 0-2 exclusive, capacity 5
	fmt.Println()
	printSlice("arryC", arryC)
	
	arryD := arryC[2:5] // here, only indexes 2-5 exclusive, capacity 5
	fmt.Println()
	printSlice("arryD", arryD)
	
}

func slice05() {
	var slce01 []int
	printSlice("slce01", slce01)
	
	if slce01 == nil {
		fmt.Println("nil! ")
		
	}
	
}

func slice04() {
	fmt.Println()
	primesArry5 := []int{2, 3, 5, 7, 11, 13}
	printSlice("primesArry5", primesArry5)
	
	// Slice the slice to give it a zero length
	fmt.Println()
	fmt.Println("Slice the slice to give it a zero length")
	primesArry5 = primesArry5[:0]
	printSlice("primesArry5", primesArry5)
	
	// Extend its length
	fmt.Println()
	fmt.Println("Extend its length")
	primesArry5 = primesArry5[:4]
	printSlice("primesArry5", primesArry5)
	
	// Drop its first two values
	fmt.Println()
	fmt.Println("Drop its first two values")
	primesArry5 = primesArry5[2:]
	printSlice("primesArry5", primesArry5)
	
}

func slice03() {
	fmt.Println()
	primesArry4 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("primesArry4: ", primesArry4)
	fmt.Println()
	primesArry4 = primesArry4[1:4]
	fmt.Println("primesArry4: ", primesArry4)
	primesArry4 = primesArry4[:2]
	fmt.Println("primesArry4: ", primesArry4)
	primesArry4 = primesArry4[1:]
	fmt.Println("primesArry4: ", primesArry4)
	fmt.Println()
	
}

func slice02() {
	fmt.Println()
	primesArry3 := []int{2, 3, 5, 7, 11, 13}
	// fmt.Println("primesArry3: ", primesArry3)
	printSlice("primesArry3 ", primesArry3)
	
	fmt.Println()
	primesArry3 = primesArry3[1:4]
	fmt.Println("execute: primesArry3 = primesArry3[1:4]")
	// fmt.Println("primesArry3: ", primesArry3)
	printSlice("primesArry3 ", primesArry3)
	fmt.Println()
	primesArry3 = primesArry3[:2]
	fmt.Println("execute: primesArry3 = primesArry3[:2]//exclusive")
	// fmt.Println("primesArry3: ", primesArry3)
	printSlice("primesArry3 ", primesArry3)
	fmt.Println()
	primesArry3 = primesArry3[1:]
	fmt.Println("execute: primesArry3 = primesArry3[1:]")
	// fmt.Println("primesArry3: ", primesArry3)
	printSlice("primesArry3 ", primesArry3)
	fmt.Println()
	
}
