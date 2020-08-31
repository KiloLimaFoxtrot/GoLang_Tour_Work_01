package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//func main() {
//	fmt.Println("Time Start: [", time.Now(), "] ")
//	fmt.Println("Hello, 世界")
//	fmt.Println("Random Num Pre-Seed 01: ", rand.Intn(30))
//	rand.Seed(16)
//	fmt.Println("Random Num Post-Seed 01: ", rand.Intn(30))
//	rand.Seed(16)
//	fmt.Println("Random Num Post-Seed 02: ", rand.Intn(30))
//	fmt.Println("Random Num Post-Seed 02: ", rand.Intn(30))
//	fmt.Println("Random Num Post-Seed 02: ", rand.Intn(30))
//	fmt.Println("Random Num Post-Seed 02: ", rand.Intn(30))
//	fmt.Println("Time End: [", time.Now(), "] ")
//
//	x := 2
//	y := 4
//	fmt.Println(add1(x, y))
//	fmt.Println(add2(x, y))
//
////	/*
////		Multiple Results:
////	*/
//	a, b := swap("hello", "world")
//	fmt.Println(a, b)
//
//	z := 6
//	c, d, e := swap2(x, y, z)
//	fmt.Println(c, d, e)
//
//}
////
/////*
////Multiple Results:
////A function can return any number of results..
////E.g. the 'swap' function below returns two strings
////*/
////
//////1.
//////'(string, string)' are the return types, unlike Java,
////// placed after function parameters
//func swap(x, y string) (string, string) {
//
//	//return x, y
//	return y, x
//}
//
////2.
////
//func swap2(x, y, z int) (int, int, int) {
//	//return x, y, z
//	//return x, z, y
//	return y, x, z
//}
////
//////Addition Long version
//func add1(x int, y int) int {
//	return x + y
//}
////
//////Addition Shorter version
////// '(int)' is the return type
//func add2(x, y int) int {
//	return x + y
//}
