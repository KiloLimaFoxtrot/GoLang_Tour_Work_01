package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println()
	fmt.Println("Defer Statements: ")

	/*
		A defer statement defers the execution of a function until the
		surrounding function returns.

		The deferred call's arguments are evaluated immediately,
		but the function call is not executed until the surrounding
		function returns

	*/
	defer01()

	/*
		Stacking defers:
		Deferred function calls are pushed onto a stack.
		When a function returns,
		its deferred calls are executed in the last-in-first-out order.

		To learn more about defer statements read this blog post.
		https://blog.golang.org/defer-panic-and-recover

	*/

	defer02()

}

func defer02() {
	fmt.Println()
	fmt.Println("Counting")

	//These Println outputs of i are first run and pushed onto a stack,
	//but deferred to return until the below Println function
	//executes, so these stacked Printlns are output in the last-in
	//-first-out order, ie in reverse.
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	//fmt.Println("done")
	fmt.Println("down")

}

func defer01() {
	var t1 = time.Now()
	defer fmt.Println("World! \n", t1)

	var t2 = time.Now()
	fmt.Println("Hello \n", t2)
}
