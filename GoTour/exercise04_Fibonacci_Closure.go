package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("**** Fibonacci Closures 01:")
	
	/*
	 Let's have some fun with functions.
	
	 Implement a fibonacci function that returns a function (
	 a closure) that returns successive fibonacci numbers (0, 1, 1,
	 2, 3, 5, ...).
	
	 fibonacci is a function that returns a function that returns
	 an int.
	*/
	
	fmt.Println()
	f := fibonacci01()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	
	// my fibonacci solution
	fmt.Println()
	fN := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fN())
	}
}

// references a solution example mentioned in the GoTour_Notes_01.
// txt file
func fibonacci01() func() int {
	fN, xN := 1, 0
	return func() int {
		fN, xN = xN, fN+xN
		return fN
	}
}

// my solution, referencing a Java example mentioned in the
// GoTour_Notes_01.txt file

// Note: in GoLang, in contrast to Java I think,
// the local variables of this function below, ie xN, yN etc..,
// seem to not reset with each iteration of this function being run
// in the main function for loop,
// possible because of the GoLang function closure technique we are
// using here..
func fibonacci() func() int {
	xN := 0
	yN := 1
	fN := 0
	cnt := 0
	
	return func() int {
		if cnt <= 0 {
			cnt++
			return xN
		} else if cnt <= 1 {
			cnt++
			return yN
		} else {
			cnt++
			// tracing the iterations, the values rotate around,
			// shift through the variables..
			fN = xN /*0, 1, 1, 2, 3, */ + yN /*1, 1, 2, 3, 5*/
			
			xN = yN /*1, 1, 2, 3, 5, */
			yN = fN /*1, 2, 3, 5, 8, */
			
			return fN /*1, 2, 3, 5, 8, ...*/
		}
	}
}
