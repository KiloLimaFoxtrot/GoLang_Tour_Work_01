package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println()
	fmt.Println("Switch Statement: ")

	Switch()

	//checks if parameter int is even or not,
	//returns a corresponding string of a color
	Switch02(5)

	/*
		Switch Statements:
		-A switch statement is a shorter way to write a sequence of
		if-else statements.
		-It runs the first case whose value is equal to the condition
		expression.
		-Go's switch is like th one in C, C++, Java, JavaScript,
		and PHP, excpet that Go only runs the selected case,
		not all the cases that follow.
		-This happens because in effect,
		the break statement that is needed at the end of each case in
		those languages is provided automatically in Go.
		-Another important difference is that Go's switch cases need
		not be constants, and the values involved need not be integers

		//@@@ Try to play with calling a function, like the above, in
		your Switch02 below
	*/

	Switch03(5)

	/*
		Switch evaluation order:
		-Switch Cases evaluate from top to bottom,
		stopping when a case succeeds

		E.g. The below will not call f() if i==0.

		Switch i {
			case 0;
			case f();


	*/

	whenSaturday()

	/*
		Note: Time in the Go playground always seems to start at
		2009-11-10 23:00:00 UTC,
		a value whose significance is left as an exercise for the
		reader..

	*/

	//??

	/*
		Switch with no condition:
		May be a better way to construct my switch03 combining int and
		bool values..
	*/

	dayGreet()

}

func dayGreet() {
	fmt.Println()
	//this will return a time object t,
	//that has time functions that can be called on it
	t := time.Now()
	fmt.Printf("Time now: %v\n", t)
	//with no switch condition, can use bool values for the cases.
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening. ")
	}
}

func specialDateTime() {
	//???

}

func whenSaturday() {
	fmt.Println()
	fmt.Println("When's Saturday? ")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today. ")
	case today + 1:
		fmt.Println("Tomorrow. ")
	case today + 2:
		fmt.Println("In two days. ")
	default:
		fmt.Println("Too far away. ")
	}
}

func Switch03(theN int) {
	fmt.Println()
	fmt.Printf("The number %v is ", theN)

	switch numEven := gratrLessEqlZero(theN % 2); numEven {
	case -1, 1:
		fmt.Println("Odd. ")
	case 0:
		fmt.Println("Even. ")
	default:
		//fmt.Printf("%v:\n", color1)
		fmt.Println("odd or even. ")
	}

}

func gratrLessEqlZero(theN int) int {

	var rslt int

	if theN > 0 {
		rslt = -1
	} else if theN < 0 {
		rslt = 1
	} else {
		rslt = 0
	}

	return rslt
}

func Switch02(theN int) {
	fmt.Println()
	fmt.Print("We prefer ")

	switch color1 := theN % 2; color1 {
	case 0:
		fmt.Println("Blue. ")
	default:
		//fmt.Printf("%v:\n", color1)
		fmt.Println("Red. ")
	}

}

func Switch() {
	fmt.Print("Go runs on ")
	//syntax to retrieve local operating system
	switch os := runtime.GOOS; os {
	case "darwin":
		//fmt.Printf("os: %v/", os)
		fmt.Println("OS X. ")
	case "linux":
		//fmt.Printf("os: %v/", os)
		fmt.Println("Linux. ")
	default: //??
		// freebsd, openbsd,
		//plan9, windows...
		fmt.Printf("%s:\n", os)

	}

	/*func Switch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X. ")
	case "linux":
		fmt.Println("Linux. ")
	default:
		// freebsd, openbsd,
		//plan9, windows...
		fmt.Printf("%s.\n", os)
	}*/
}
