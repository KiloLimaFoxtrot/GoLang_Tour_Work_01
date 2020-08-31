package main

import "fmt"

func main() {

	// For Loop 01
	/*
		The For Loop has three components separated by semicolons:
		1. init (initial?) statement:
		-Executed before the first iteration
		-Often is a short variable declaration,
		which will only be visible within the scope of the For Loop

		2. condition expression:
		-A boolean expression that serves as a condition check
		-Evaluated before every iteration,
		if conditions do not meet the condition expressions criteria,
		then the expression returns false and the program will exit the
		For Loop

		3. post statement:
		-Executed at the end of every iteration

	*/

	sum := 0
	for i := 0; i < 20; i++ {
		sum += i
	}
	fmt.Println()
	fmt.Println("For Loop 01 ")
	fmt.Println("sum value: ", sum)

	// For Loop 02
	sum2 := 1
	for ; sum2 < 1000; {
		sum2 += sum2
	}

	fmt.Println()
	fmt.Println("For Loop 02 ")
	fmt.Println("Sum2 value: ", sum2)

	// For Loop 03; The While Loop
	/*
		For Loop / While Loop

		With Go, the For Loop is the While Loop.
		To use the For Loop as a while loop,
		drop the semicolons after the for..

	*/

	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println()
	fmt.Println("For Loop 03 ")
	fmt.Println("For Loop as While Loop ")
	fmt.Printf("Sum3 value: %v\n", sum3)

	// For Loop 03; The Infinite Loop
	// This loop is commented out because the IDE detects the
	//infinite loop which will cause an error.
	/*
		for {
			//no conditions that will
		}
	*/

}
