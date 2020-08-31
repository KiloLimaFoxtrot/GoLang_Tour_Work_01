package main

//
// import (
// 	"fmt"
// )
//
// func main() {
// 	fmt.Println()
// 	fmt.Println("**** Fibonacci Closures 01:")
//
// 	/*
// 	 Let's have some fun with functions.
//
// 	 Implement a fibonacci function that returns a function (
// 	 a closure) that returns successive fibonacci numbers (0, 1, 1,
// 	 2, 3, 5, ...).
//
// 	 fibonacci is a function that returns a function that returns
// 	 an int.
// 	*/
//
// 	fmt.Println()
// 	f := fibonacci01()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(f())
// 	}
//
// 	// my fibonacci test
// 	fmt.Println()
// 	fT := fibonacci()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(fT())
// 	}
// }
//
// func fibonacci01() func() int {
// 	fN, xN := 1, 0
// 	return func() int {
// 		fN, xN = xN, fN+xN
// 		return fN
// 	}
// }
//
// // my test, referencing a Java example mentioned in the
// // GoTour_Notes_01.txt file
// func fibonacci() func() int {
// 	xN := 0
// 	yN := 1
// 	fN := 0
// 	cnt := 0
//
// 	// return func(i int) int {
// 	// 	if i <= 0 {
// 	// 		return xN
// 	// 	} else if i <= 1 {
// 	// 		return yN
// 	// 	} else {
// 	// 		fN = xN /*0, 1*/ + yN /*1, 1*/
// 	// 		xN = yN              /*1, 1*/
// 	// 		yN = fN              /*1, 2*/
// 	// 		return fN /*1, 2 ...*/
// 	// 	}
// 	// }
//
// 	return func() int {
// 		if cnt <= 0 {
// 			cnt++
// 			return xN
// 		} else if cnt <= 1 {
// 			cnt++
// 			return yN
// 		} else {
// 			cnt++
// 			fN = xN /*0, 1*/ + yN /*1, 1*/
// 			xN = yN               /*1, 1*/
// 			yN = fN               /*1, 2*/
// 			return fN             /*1, 2 ...*/
// 		}
// 	}
// }
//
// // the Fibonacci is, a function that returns a function that returns
// // an int
// // func fibonacci01() func() int {
// // 	F := 2
// // 	return func() int {
// // 		if F < 2 {
// // 			return F
// // 		} else {
// // 			F += (F - 1) + (F - 2)
// // 			return F
// // 		}
// // 	}
// // }
//
// // func fibonacci01() func() int {
// // 	F := 0
// // 	return func() int {
// // 		if F < 2 {
// // 			F++
// // 		} else {
// // 			F += (F - 1) + (F - 2)
// // 		}
// // 		return F
// // 	}
// // }
