package main

import "fmt"

func main() {

	fmt.Println()
	fmt.Println("Slices: ")
	fmt.Println("(Arrays cntd...) ")

	//An array has a fixed size. A slice, on the other had,
	//is a dynamically-sized,
	//flexible vew into the elements of an array. In practice,
	//slices are much more common than arrays.

	//The []T is a slice with a type T.

	//A Slice is formed by specifying two indices,
	//a low and high bound, separated by a colon:

	//e.g.
	//  a[low: high] //excludes last element

	//This selects a half-open range which includes the first
	//element, but excludes the last one

	//The following expression creates a slice which includes
	//elements 1 through 3 of a:

	//e.g.
	// 	a[1: 4]

	//Slices are like references to arrays
	//A slice does not store any data,
	//it just describes a section of an underlying array.

	//Changing the elements of a slice modifies the corresponding
	//elements of its underlying array

	//Other slices that share the same underlying array will see
	//those changes

	//Slice Literals:

	//A slice literal is like an array literal, but without the length

	//E.g. this is an array literal

	//  [3]bool{true, tue, false}

	//And this creates the same array as above,
	//then builds a slice that references it:

	//A slice literal?
	//  []bool{true, true, false}

	slice01()

}

func slice01() {

	primesArry2 := [6]int{2, 3, 5, 7, 11, 13}

	namesArry01 := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	var s []int = primesArry2[1:4] //this will create a slice of
	// primesArry2, ie a sub-array of primesArry2

	fmt.Println()
	fmt.Printf("Slice s: %v \n", s)

	fmt.Println()
	fmt.Printf("namesArry01: %v \n", namesArry01)

	//slices of the namesArray01
	namesSlice01 := namesArry01[0:2]
	namesSlice02 := namesArry01[1:3]
	fmt.Println()
	fmt.Printf("namesSlice01: %v \n", namesSlice01)
	fmt.Printf("namesSlice02: %v \n", namesSlice02)

	//changing the elements of a slice modifies the corresponding
	//elements of its underlying array.

	namesSlice02[0] = "XXX"
	fmt.Println()
	fmt.Println("executing -> namesSlice02[0] = \"XXX\"")

	fmt.Println()
	fmt.Printf("namesSlice01: %v \n", namesSlice01)
	fmt.Printf("namesSlice02: %v \n", namesSlice02)
	fmt.Printf("namesArry01: %v \n", namesArry01)

	fmt.Println()
	var arryLtrl = [3]bool{true, true, false}
	fmt.Println("Array Literal -> var arryLtrl = [3]bool{true, true, " +
		"false}")
	fmt.Printf("Array Literal: %v \n", arryLtrl)

	fmt.Println()
	var arrySlceLtrl = []bool{true, true, false}
	fmt.Println("Array Slice Literal -> var arrySlceLtrl = []bool{true, " +
		"true, false}")
	fmt.Printf("Array Slice Literal: %v \n", arrySlceLtrl)

	arrySlceLtrlQ := []int{2, 3, 5, 7, 11, 13}
	fmt.Println()
	fmt.Printf("Array Slice Literal -> arrySlceLtrlQ := []int{2, 3, 5, 7, 11, 13}")
	fmt.Printf("q: %v \n", arrySlceLtrlQ)

	arrySlceLtrlR := []bool{true, false, true, true, false, true}
	fmt.Println()
	fmt.Printf("Array Slice Literal -> arrySlceLtrlR := []bool{true, " +
		"false, true, true, false, true}")
	fmt.Printf("q: %v \n", arrySlceLtrlR)

	//An array that is of the type struct,
	//so can create lots of versions with differing values
	structS := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println()
	fmt.Printf("structS %v \n", structS)

}
