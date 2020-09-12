package main

import "fmt"

func main() {
	fmt.Println()
	fmt.Println("Arrays: ")

	//The type [n]T is an array of n values of type T..

	//The expression

	//	var a [10]int

	//declares a variable 'a' as an array of ten integers

	//An arrays length is part of its type,
	//so arrays cannot be resized. This seems limiting,
	//but don't worry; Go provides a convenient way of working with
	//arrays.

	arrays01()

}

func arrays01() {
	var arry01 [2]string
	arry01[0] = "Hello"
	arry01[1] = "World"

	fmt.Println()
	fmt.Printf("arry01[0]: %v, arry01[1]: %v \n", arry01[0], arry01[1])
	fmt.Printf("arry01: %v \n", arry01)

	primesArry := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println()
	fmt.Printf("primesArry %v \n", primesArry)

}
