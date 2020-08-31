package main

import "fmt"

var (
	powrsArry = []int{1, 2, 4, 8, 16, 32, 64, 128}
)

func main() {
	fmt.Println()
	fmt.Println("Range: ")

	// The range form of the for loop iterated over a slice or map

	// When randing over a slice,
	//two values are returned for each iteration.

	// The first is the index,
	//and the second is a copy of the element (value?) at that index

	fmt.Println()
	fmt.Println("  **FUNCTION range01: ")
	range01()

}

func range01() {
	fmt.Println()

	//'range' form, ie GoLang function?
	for i, v := range powrsArry {
		fmt.Printf("2**%d = %d \n", i, v)
	}

	fmt.Println()
	for i, val := range powrsArry {
		fmt.Printf("Iteration: %v, Index: %v, Value: %v \n", i+1, i,
			val)
	}

	fmt.Println()
	for i, val := range powrsArry {
		fmt.Printf("Iteration: %d, Index: %d, Value: %d \n", i+1, i,
			val)
	}

	// You can skip the index or value by assigning to '_'
	fmt.Println()
	for i, _ := range powrsArry {
		fmt.Printf("Iteration: %v, Index: %v \n", i+1, i)
	}

	// You can skip the index or value by assigning to '_'
	fmt.Println()
	for _, val := range powrsArry {
		fmt.Printf("Value: %v \n", val)
	}

}
